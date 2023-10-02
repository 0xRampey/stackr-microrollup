package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func backfillSubmissions(client *ethclient.Client, l1Contract common.Address) []common.Hash {
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

func backfillRegistrations(client *ethclient.Client, l1Contract common.Address) []common.Address {
	// Calculate the signature of the event
	eventSignature := "AppRegistered(address)"
	hash := crypto.Keccak256Hash([]byte(eventSignature))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{l1Contract},
		Topics:    [][]common.Hash{{hash}},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	addresses := []common.Address{}
	// Handle logs from the past (backfill)
	for _, vLog := range logs {
		addressBytes := vLog.Topics[1].Bytes()
		address := common.BytesToAddress(addressBytes)
		addresses = append(addresses, address)
	}
	return addresses
}

func subscribeToSubmissions(client *ethclient.Client, l1Contract common.Address) (ethereum.Subscription, chan ethtypes.Log) {
	// Calculate the signature of the event
	eventSignature := "BatchSubmitted(bytes32)"
	hash := crypto.Keccak256Hash([]byte(eventSignature))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{l1Contract},
		Topics:    [][]common.Hash{{hash}},
	}
	// Subscribe to new logs
	logsCh := make(chan ethtypes.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logsCh)
	if err != nil {
		log.Fatal(err)
	}
	return sub, logsCh
}
