package main

import (
	"bytes"
	"encoding/json"
	"github.com/ardanlabs/blockchain/foundation/blockchain/database"
	"log"
	"net/http"
)

func main() {
	err := sendTransaction()

	if err != nil {
		log.Fatalf("error sending transaction: %s", err)
	}
}

func sendTransaction() error {

	tx, err := database.NewTx(1, 1, "bill", "nathan", 1000000, 0, nil)

	if err != nil {
		return err
	}

	data, err := json.Marshal(tx)

	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080/v1/tx/submit",
		"application/json", bytes.NewBuffer(data),
	)

	if err != nil {
		return err
	}

	return resp.Body.Close()
}
