package main

import (
	"log"
	"stackr-mvp/types"
)

func main() {
	// Send batches to aggregator through channels instead of RPC
	batchChannel := make(chan types.Batch)
	app := RollApp{}
	app.InitState()
	go app.InitServer(batchChannel)
	go app.subscribeToDeposits()
	go app.subscribeToSubmissions()

	agg := Aggregator{}
	agg.Init()
	go agg.subscribeToSubmissions()

	log.Println("Aggregator waiting for batches.....")

	for batch := range batchChannel {
		agg.submitBatch(batch)
	}
}
