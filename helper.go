package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func backfillDeposits(client *ethclient.Client, l1Contract common.Address) []common.Hash {
	// Calculate the signature of the event
	eventSignature := "BatchSubmitted(bytes32)"
	hash := crypto.Keccak256Hash([]byte(eventSignature))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{l1Contract},
		Topics:    [][]common.Hash{{hash}},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	batchHashes := []common.Hash{}
	// Handle logs from the past (backfill)
	for _, vLog := range logs {
		batchHashes = append(batchHashes, common.Hash(vLog.Topics[1].Bytes()))
		log.Println("Backfill Submission Event!")
	}
	return batchHashes
}
