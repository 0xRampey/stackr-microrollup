package main

import (
	"log"
	"stackr-mvp/types"
)

func main() {
	// Send batches to aggregator through channels instead of RPC
	batchChannel := make(chan types.Batch)
	app := RollApp{}
	go app.Init(batchChannel)

	agg := Aggregator{}
	agg.Init()

	log.Println("Aggregator waiting for batches.....")

	for batch := range batchChannel {
		agg.submitBatch(batch)
	}
}
