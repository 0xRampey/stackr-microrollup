package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"stackr-mvp/types"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Load a user's private key
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	message := map[string]interface{}{
		"action": "add_todo",
		"index":  0,
	}
	messageBytes, _ := json.Marshal(message)
	messageStr := string(messageBytes)

	hash := crypto.Keccak256Hash([]byte(messageStr))
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	request := &types.Tx{
		Message:   messageStr,
		Signature: hexutil.Encode(signature),
	}
	requestBytes, _ := json.Marshal(request)

	url := "http://localhost:8080/tx" // replace with your actual server URL
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBytes))
	if err != nil {
		fmt.Printf("Error while sending request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Server returned error: %s\n", resp.Status)
		return
	}

	fmt.Println("Message sent successfully!")
}
