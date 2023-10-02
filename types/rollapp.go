package types

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/cbergoon/merkletree"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

type Batch struct {
	Header BatchHeader
	TxList []Tx
}

func (b Batch) GetSignature(pk *ecdsa.PrivateKey) string {
	// Convert the batch structure to bytes
	batchBytes, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate the hash of the batch bytes
	batchHash := crypto.Keccak256Hash(batchBytes)

	// Sign the batch hash using the private key
	signature, err := crypto.Sign(batchHash.Bytes(), pk)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the signature to hex string
	signatureHex := hexutil.Encode(signature)

	return signatureHex
}

func (b Batch) GetSigneeAddress(signatureHex string) common.Address {
	// Convert the batch structure to bytes
	batchBytes, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate the hash of the batch bytes
	batchHash := crypto.Keccak256Hash(batchBytes)

	// Decode the signature from hex string to bytes
	signatureBytes, err := hexutil.Decode(signatureHex)
	if err != nil {
		log.Fatal(err)
	}

	// Verify the signature using the public key
	publicKey, err := crypto.Ecrecover(batchHash.Bytes(), signatureBytes)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the public key to ECDSA format
	publicKeyECDSA, err := crypto.UnmarshalPubkey(publicKey)
	if err != nil {
		log.Fatal(err)
	}

	// Get the address from the public key
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return address
}

type SignedBatch struct {
	Batch
	Signature string
}

type BatchHeader struct {
	PrevHash  common.Hash // Hash of previous batch's header
	StateRoot common.Hash // Hash of the state tree root
	TxRoot    common.Hash // Hash of the tx tree root
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
	h := sha3.NewLegacyKeccak256()
	if _, err := h.Write([]byte(t.Signature)); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func (t Tx) Equals(other merkletree.Content) (bool, error) {
	return bytes.Equal([]byte(t.Signature), []byte(other.(Tx).Signature)), nil
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
