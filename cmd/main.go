package main

import (
	"fmt"
	"time"

	"github.com/pokt-foundation/tx-client/txclient"
)

func main() {
	client, _ := txclient.NewTXClient(txclient.Config{
		BaseURL: "http://localhost:8080",
		APIKey:  "test_api_key",
		Version: txclient.V0,
		Timeout: 5 * time.Second,
	})

	sr, err := client.GetServiceRecord(1)
	if err != nil {
		panic("err")
	}
	fmt.Printf("%v\n", sr)
}
