package main

import (
	"fmt"

	"github.com/cbergoon/merkletree"
	"github.com/gin-gonic/gin"
)

type RollApp struct {
	server *gin.Engine
	db     []merkletree.Content
	tree   *merkletree.MerkleTree
	// list   []Tx
}

func (r *RollApp) Init() {
	// Backfill past events
	// Calculate state tree
	// Start server
	r.server = gin.Default()
	r.server.GET("/tx", r.handleTx)
	fmt.Println("RollApp initialized")
	r.server.Run(":8080")
}

func (r *RollApp) handleTx(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "pong",
	})
	// Validate tx
	// Make sure user exists
	// Make sure nonce is correct
	// Make sure tx is signed
	// Update and create new state tree
	// Add tx to list
	// Submit batch to aggregator if we have enough txs
}
