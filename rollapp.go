package main

import (
	"fmt"
	"stackr-mvp/types"

	"github.com/cbergoon/merkletree"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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
	r.server.POST("/tx", r.handleTx)
	fmt.Println("RollApp initialized")
	r.server.Run(":8080")
}

func (r *RollApp) handleTx(c *gin.Context) {
	var tx types.Tx
	if err := c.ShouldBindJSON(&tx); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	sigBytes, _ := hexutil.Decode(tx.Signature)

	hash := crypto.Keccak256Hash([]byte(tx.Message))
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
	c.JSON(200, gin.H{"status": "State transition successful"})
	// Validate tx
	// Make sure user exists
	// Make sure nonce is correct
	// Make sure tx is signed
	// Update and create new state tree
	// Add tx to list
	// Submit batch to aggregator if we have enough txs
}
