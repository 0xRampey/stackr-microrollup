package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	messageBytes, _ := json.Marshal(tx.Message)
	hash := crypto.Keccak256Hash([]byte(messageBytes))
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

	for idx, content := range r.db {
		userTodos := content.(UserTodos)
		if userTodos.account == address.Hex() {
			// Found user
			log.Println("Found user!")
			// Try to update user's todos
			newUserTodo, err := r.updateState(userTodos, tx.Message)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

			// Update db and recalculate merkle tree
			r.db[idx] = newUserTodo
			t, err := merkletree.NewTree(r.db)
			if err != nil {
				log.Fatal(err)
			}
			r.tree = t

			// Update tx list

			c.JSON(200, gin.H{"status": "State transition successful"})
		}
	}

	c.JSON(200, gin.H{"status": "State transition failed! Could not find user!"})
	log.Println("State transition failed! Could not find user!")
	// Validate tx
	// Make sure user exists
	// Make sure nonce is correct
	// Make sure tx is signed

	// Add tx to list
	// Submit batch to aggregator if we have enough txs
}

func (r *RollApp) updateState(userTodos UserTodos, m types.Message) (UserTodos, error) {
	switch m.Action {
	case "add_todo":
		userTodos.todos = append(userTodos.todos, m.Content)
	case "mark_done":
		if m.Index < len(userTodos.todos) && m.Index >= 0 {
			userTodos.todos[m.Index] = userTodos.todos[len(userTodos.todos)-1]
			userTodos.todos = userTodos.todos[:len(userTodos.todos)-1]
		} else {
			return UserTodos{}, fmt.Errorf("Invalid index: %d", m.Index)
		}
	default:
		return UserTodos{}, fmt.Errorf("Invalid action: %s", m.Action)
	}

	return userTodos, nil
}
