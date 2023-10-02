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
	registeredApps   []common.Address
	latestHeaderHash common.Hash
	ethClient        *ethclient.Client
	l1Contract       common.Address
}

func (a *Aggregator) Init() {
	// Init eth client
	client, err := ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	a.ethClient = client

	a.l1Contract = common.HexToAddress("0x8464135c8F25Da09e49BC8782676a84730C318bC")
	// Backfill past events (or just get the last 'settled' batch header)
	a.backfill()
}

func (a *Aggregator) submitBatch(b types.SignedBatch) {
	// Verify signature from registered app
	if !a.verifyAppRegistered(b) {
		log.Println("App not registered in", a.registeredApps)
		return
	}
	batch := b.Batch
	// If this is the genesis batch, ignore batch ordering check
	if a.latestHeaderHash == (common.Hash{}) {
		a.latestHeaderHash = batch.Header.CalculateHash() // TODO: This should be set only after L1 confirmation
	} else {
		// Verify batches are submitted in order
		if !bytes.Equal(batch.Header.PrevHash.Bytes(), a.latestHeaderHash.Bytes()) {
			log.Println(batch.Header.PrevHash.Bytes(), a.latestHeaderHash.Bytes())
			log.Println("Batch submitted out of order")
			return
		}
	}
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
		StateRoot: [32]byte(batch.Header.StateRoot.Bytes()),
		TxRoot:    [32]byte(batch.Header.TxRoot.Bytes()),
	}

	settlementTxs := make([]settlement.SettlementTx, len(batch.TxList))
	for i, tx := range batch.TxList {
		settlementTxs[i] = settlement.SettlementTx{
			Signature: []byte(tx.Signature),
		}
	}

	privateKey, err := crypto.HexToECDSA("59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d")
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

	log.Printf("batch submitted: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

}

func (a *Aggregator) backfill() {
	batchHashes := backfillSubmissions(a.ethClient, a.l1Contract)
	if len(batchHashes) > 0 {
		a.latestHeaderHash = batchHashes[len(batchHashes)-1]
	}

	a.registeredApps = backfillRegistrations(a.ethClient, a.l1Contract)
}

func (a *Aggregator) subscribeToSubmissions() {
	sub, logsCh := subscribeToSubmissions(a.ethClient, a.l1Contract)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logsCh:
			log.Println("Aggregator received batch confirmation!")
			a.latestHeaderHash = common.Hash(vLog.Topics[1].Bytes())
		}
	}
}

func (a *Aggregator) verifyAppRegistered(b types.SignedBatch) bool {
	sig := b.Signature

	signeeAddress := b.Batch.GetSigneeAddress(sig)
	log.Println("Submitted app's address: ", signeeAddress.Hex())

	for _, app := range a.registeredApps {
		if app == signeeAddress {
			return true
		}
	}
	return false
}
