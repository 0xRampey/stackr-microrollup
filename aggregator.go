package main

import (
	"log"
	"stackr-mvp/types"

	"github.com/ethereum/go-ethereum/common"
)

type Aggregator struct {
	registered_apps []common.Address
	// last_batch      types.BatchHeader
}

func (a *Aggregator) Init() {
	// Get registered apps from L1 contract
	// Backfill past events (or just get the last 'settled' batch header)
}

func (a *Aggregator) submitBatch(b types.Batch) {
	log.Println("Submitting batch", b)
	// Verify batch
	// Verify signature
	// Verify merkle root
	// Verify state transition
	// Submit batch to L1 contract
}
