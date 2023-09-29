package types

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

	"github.com/cbergoon/merkletree"
)

type Batch struct {
	Tx_list []Tx
}

type Tx struct { // Let's keep it simple for now
	Message   Message `json:"message"`
	Signature string  `json:"signature"`
}

type Message struct {
	Action  string `json:"action"`
	Content string `json:"content"`
	Index   int    `json:"index"`
}

type UserTodos struct {
	Account string
	Nonce   int
	Todos   []string
}

func (ut UserTodos) CalculateHash() ([]byte, error) {
	h := sha256.New()
	// HACK: This is not a good way to calculate the hash. (Use RLP encoding in prod)
	input := ut.Account + fmt.Sprint(ut.Nonce) + strings.Join(ut.Todos, ",")
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
	return ut.Account == otherUT.Account, nil
}
