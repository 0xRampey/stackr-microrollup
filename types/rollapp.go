package types

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/cbergoon/merkletree"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Batch struct {
	Header  BatchHeader
	Tx_list []Tx
}

type BatchHeader struct {
	PrevHash  common.Hash // Hash of previous batch's header
	StateRoot *merkletree.Node
	TxRoot    *merkletree.Node
}

func (bh BatchHeader) CalculateHash() common.Hash {
	combined := append(bh.PrevHash.Bytes(), []byte(bh.StateRoot.String())...)
	return common.Hash(crypto.Keccak256Hash(combined))
}

type Tx struct { // Let's keep it simple for now
	Message   Message `json:"message"`
	Signature string  `json:"signature"`
}

func (t Tx) CalculateHash() ([]byte, error) {
	return crypto.Keccak256Hash([]byte(t.Signature)).Bytes(), nil
}

func (t Tx) Equals(other merkletree.Content) (bool, error) {
	otherTx, ok := other.(Tx)
	if !ok {
		return false, errors.New("value is not of type Tx")
	}
	return t.Signature == otherTx.Signature, nil
}

type Message struct {
	Action  string `json:"action"`
	Content string `json:"content"`
	Index   uint   `json:"index"`
}

type UserTodos struct {
	Account string
	Nonce   int
	Todos   []string
	Balance *big.Int
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

// Define a struct for the Deposit event using tags to associate event field names to struct fields
type DepositEvent struct {
	User   common.Address
	Amount *big.Int
}
