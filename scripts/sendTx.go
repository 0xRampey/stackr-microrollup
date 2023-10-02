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
	privateKey, err := crypto.HexToECDSA("dbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97")
	if err != nil {
		log.Fatal(err)
	}

	message := types.Message{
		Action:  "add_todo",
		Index:   0,
		Content: "get milk",
	}
	messageBytes, _ := json.Marshal(message)

	hash := crypto.Keccak256Hash(messageBytes)
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	request := &types.Tx{
		Message:   message,
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
	fmt.Println("Response: ", resp.Status, resp.Body)
}
