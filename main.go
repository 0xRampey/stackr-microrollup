package main

import "stackr-mvp/types"

func main() {
	agg := Aggregator{}
	agg.Init()

	batchChannel := make(chan types.Batch)

	app := RollApp{}
	app.Init(&batchChannel)

	for batch := range batchChannel {
		agg.submitBatch(batch)
	}
}
