package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

	"github.com/cbergoon/merkletree"
)

func main() {
	//Build list of Content to build tree
	var list []merkletree.Content
	list = append(list, UserTodos{account: "Alice", nonce: 1, todos: []string{"a", "b", "c"}})
	list = append(list, UserTodos{account: "Bob", nonce: 2, todos: []string{"a", "b", "c"}})

	// //Get the Merkle Root of the tree
	// mr := t.MerkleRoot()
	// log.Println(mr)

	// //Verify the entire tree (hashes for each node) is valid
	// vt, err := t.VerifyTree()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Verify Tree: ", vt)

	app := RollApp{}
	app.Init()
}

type UserTodos struct {
	account string
	nonce   int
	todos   []string
}

func (ut UserTodos) CalculateHash() ([]byte, error) {
	h := sha256.New()
	// HACK: This is not a good way to calculate the hash. (Use RLP encoding in prod)
	input := ut.account + fmt.Sprint(ut.nonce) + strings.Join(ut.todos, ",")
	if _, err := h.Write([]byte(input)); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// Equals tests for equality of two Contents
func (ut UserTodos) Equals(other merkletree.Content) (bool, error) {
	otherUT, ok := other.(UserTodos)
	if !ok {
		return false, errors.New("value is not of type UserTodos")
	}
	return ut.account == otherUT.account, nil
}
