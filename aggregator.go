package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	settlement "stackr-mvp/contracts"
	"stackr-mvp/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Aggregator struct {
	registered_apps []common.Address
	latest_header   types.BatchHeader
	ethClient       *ethclient.Client
	l1Contract      common.Address
}

func (a *Aggregator) Init() {
	// Init eth client
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	a.ethClient = client

	a.l1Contract = common.HexToAddress("0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6")
	// Get registered apps from L1 contract
	// Backfill past events (or just get the last 'settled' batch header)
}

func (a *Aggregator) submitBatch(batch types.Batch) {
	// TODO: Verify signature

	// If this is the genesis batch, set the latest header
	if a.latest_header == (types.BatchHeader{}) {
		a.latest_header = batch.Header
	} else {
		// Verify batches are submitted in order
		if !bytes.Equal(batch.Header.PrevHash.Bytes(), a.latest_header.CalculateHash().Bytes()) {
			log.Println("Batch submitted out of order")
			return
		}
	}

	// TODO: Verify state transition?
	// tODO: Sign the batch
	// Submit batch to L1 contract
	a.submitToL1(batch)

}

func (a *Aggregator) submitToL1(batch types.Batch) {
	contract, err := settlement.NewSettlement(a.l1Contract, a.ethClient)
	if err != nil {
		log.Fatal(err)
	}

	batchHeader := settlement.SettlementBatchHeader{
		PrevHash:  batch.Header.PrevHash,
		StateRoot: [32]byte(batch.Header.StateRoot.Hash),
		TxRoot:    [32]byte(batch.Header.TxRoot.Hash),
	}

	settlementTxs := make([]settlement.SettlementTx, len(batch.Tx_list))
	for i, tx := range batch.Tx_list {
		settlementTxs[i] = settlement.SettlementTx{
			Signature: []byte(tx.Signature),
		}
	}

	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Println(fromAddress.Hex()) // 0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1
	nonce, err := a.ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := a.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := contract.SubmitBatch(auth, batchHeader, settlementTxs)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

}
