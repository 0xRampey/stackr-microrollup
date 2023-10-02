package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	settlement "stackr-mvp/contracts"
	"stackr-mvp/types"

	"github.com/cbergoon/merkletree"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type RollApp struct {
	server           *gin.Engine
	db               []merkletree.Content
	tree             *merkletree.MerkleTree
	tx_list          []types.Tx
	latestHeaderHash common.Hash            // Convenience parameter
	batch_channel    chan types.SignedBatch // Send batches to aggregator instead of RPC
	ethClient        *ethclient.Client
	l1Contract       common.Address
	privateKey       *ecdsa.PrivateKey // Unsafe, I know! But this is just a demo. In prod, it'd be a signer loaded from a gitignored .env file
}

func (r *RollApp) InitState() {
	log.Println("Initializing RollApp state.....")
	// Init eth client
	client, err := ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	r.ethClient = client

	r.l1Contract = common.HexToAddress("0x8464135c8F25Da09e49BC8782676a84730C318bC")

	// Backfill past events
	r.backfill()

	// Calculate state tree
	t, err := merkletree.NewTree(r.db)
	if err != nil {
		log.Fatal(err)
	}
	r.tree = t

	// Load a user's private key
	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		log.Fatal(err)
	}
	r.privateKey = privateKey

	// Register app on L1 contract
	r.registerApp()
}

func (r *RollApp) InitServer(c chan types.SignedBatch) {
	// Start server
	r.server = gin.Default()
	r.server.POST("/tx", r.handleTx)
	log.Println("RollApp ready to receive txs.....")

	// Initialise batch channel
	r.batch_channel = c
	r.server.Run(":8080")
}

func (r *RollApp) backfill() {

	// Calculate the signature of the event
	eventSignature := "Deposit(address,uint256)"
	hash := crypto.Keccak256Hash([]byte(eventSignature))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{r.l1Contract},
		Topics:    [][]common.Hash{{hash}},
	}

	logs, err := r.ethClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	// Handle logs from the past (backfill)
	for _, vLog := range logs {
		deposit := types.DepositEvent{
			User:   common.BytesToAddress(vLog.Topics[1].Bytes()),
			Amount: new(big.Int).SetBytes(vLog.Data),
		}
		log.Printf("Backfill Deposit Event: %+v\n", deposit)
		r.db = append(r.db, types.UserTodos{
			Account: deposit.User.Hex(),
			Nonce:   0,
			Todos:   []string{},
			Balance: deposit.Amount,
		})
	}

	// Also backfill batch submissions
	batchHashes := backfillSubmissions(r.ethClient, r.l1Contract)
	if len(batchHashes) > 0 {
		r.latestHeaderHash = batchHashes[len(batchHashes)-1]
	}
}

func (r *RollApp) subscribeToDeposits() {
	// Calculate the signature of the event
	eventSignature := "Deposit(address,uint256)"
	hash := crypto.Keccak256Hash([]byte(eventSignature))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{r.l1Contract},
		Topics:    [][]common.Hash{{hash}},
	}
	// Subscribe to new logs
	logsCh := make(chan ethtypes.Log)
	sub, err := r.ethClient.SubscribeFilterLogs(context.Background(), query, logsCh)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("RollApp subscribed to L1.....")

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logsCh:
			deposit := types.DepositEvent{
				User:   common.BytesToAddress(vLog.Topics[1].Bytes()),
				Amount: new(big.Int).SetBytes(vLog.Data),
			}
			log.Printf("New Deposit Event: %+v\n", deposit)
			r.db = append(r.db, types.UserTodos{
				Account: deposit.User.Hex(),
				Nonce:   0,
				Todos:   []string{},
				Balance: deposit.Amount,
			})
		}
	}

}

func (r *RollApp) subscribeToSubmissions() {
	sub, logsCh := subscribeToSubmissions(r.ethClient, r.l1Contract)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logsCh:
			log.Println("RollApp received batch confirmation!")
			r.latestHeaderHash = common.Hash(vLog.Topics[1].Bytes())
		}
	}
}

func (r *RollApp) handleTx(c *gin.Context) {
	var tx types.Tx
	if err := c.ShouldBindJSON(&tx); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	sigBytes, _ := hexutil.Decode(tx.Signature)

	messageBytes, _ := json.Marshal(tx.Message)
	hash := crypto.Keccak256Hash([]byte(messageBytes))
	pubKey, err := crypto.Ecrecover(hash.Bytes(), sigBytes)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": "Failed to recover public key"})
		return
	}
	pubKeyECDSA, err := crypto.UnmarshalPubkey(pubKey)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to unmarshal public key"})
		return
	}
	address := crypto.PubkeyToAddress(*pubKeyECDSA)

	// Perform state transition here, for example, updating balances, and respond
	// accordingly
	fmt.Printf("State transition requested by %s\n", address.Hex())

	for idx, content := range r.db {
		userTodos := content.(types.UserTodos)
		if userTodos.Account == address.Hex() {
			// Found user
			log.Println("Found user!")
			// Try to update user's todos
			newUserTodo, err := r.updateState(userTodos, tx.Message)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

			// Update db and recalculate merkle tree
			r.db[idx] = newUserTodo
			t, err := merkletree.NewTree(r.db)
			if err != nil {
				log.Fatal(err)
			}
			r.tree = t

			// Update tx list
			r.updateTxList(tx)

			log.Printf("User's new state: %+v\n", newUserTodo)

			c.JSON(200, gin.H{"status": "State transition successful"})
			return
		}
	}

	c.JSON(200, gin.H{"status": "State transition failed! Could not find user!"})
	log.Println("State transition failed! Could not find user!")

}

func (r *RollApp) updateState(userTodos types.UserTodos, m types.Message) (types.UserTodos, error) {
	switch m.Action {
	case "add_todo":
		userTodos.Todos = append(userTodos.Todos, m.Content)
		userTodos.Balance.Sub(userTodos.Balance, big.NewInt(100000)) // Deduct for gas fees
		// TODO: Send gas fees to App/Aggregator
		userTodos.Nonce++
	case "mark_done":
		if int(m.Index) < len(userTodos.Todos) {
			userTodos.Todos[m.Index] = userTodos.Todos[len(userTodos.Todos)-1]
			userTodos.Todos = userTodos.Todos[:len(userTodos.Todos)-1]
			userTodos.Balance.Sub(userTodos.Balance, big.NewInt(100000)) // Deduct for gas fees
			// TODO: Send gas fees to App/Aggregator
			userTodos.Nonce++
		} else {
			return types.UserTodos{}, fmt.Errorf("Invalid index: %d", m.Index)
		}
	default:
		return types.UserTodos{}, fmt.Errorf("Invalid action: %s", m.Action)
	}

	return userTodos, nil
}

func (r *RollApp) updateTxList(tx types.Tx) {
	r.tx_list = append(r.tx_list, tx)
	if len(r.tx_list) > 3 { // Batch every x txs (2 here for testing purposes)
		// Submit batch to aggregator
		log.Println("Submitting batch to aggregator.....")
		prevHash := common.Hash(crypto.Keccak256Hash([]byte("Genesis Batch")))
		// Genesis batch
		if r.latestHeaderHash == (common.Hash{}) {
		} else {
			prevHash = r.latestHeaderHash
		}

		blocks := make([]merkletree.Content, 0, len(r.tx_list))
		for _, tx := range r.tx_list {
			blocks = append(blocks, tx)
		}
		txTree, err := merkletree.NewTree(blocks)
		if err != nil {
			log.Fatal(err)
		}

		header := types.BatchHeader{
			PrevHash:  prevHash,
			StateRoot: common.Hash(r.tree.MerkleRoot()),
			TxRoot:    common.Hash(txTree.MerkleRoot()),
		}
		batch := types.Batch{
			Header: header,
			TxList: r.tx_list,
		}
		// Sign batch

		r.batch_channel <- types.SignedBatch{
			Batch:     batch,
			Signature: batch.GetSignature(r.privateKey),
		}
		// Flush txs
		r.tx_list = []types.Tx{}
	}
}

func (r *RollApp) registerApp() {
	contract, err := settlement.NewSettlement(r.l1Contract, r.ethClient)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := r.privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Println(fromAddress.Hex()) // 0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1
	nonce, err := r.ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := r.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(r.privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := contract.RegisterApp(auth)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("App registration sent: %s", tx.Hash().Hex())
}
