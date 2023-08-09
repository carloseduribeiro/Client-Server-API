package main

import (
	"context"
	"encoding/json"
	"github.com/carloseduribeiro/Client-Server-API/client/client"
	"log"
	"os"
)

func main() {
	e, err := client.GetExchange(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	encoder := json.NewEncoder(os.Stdout)
	if err = encoder.Encode(e); err != nil {
		panic(err)
	}
}
