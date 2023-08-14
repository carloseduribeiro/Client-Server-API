package main

import (
	"context"
	"encoding/json"
	"github.com/carloseduribeiro/Client-Server-API/client/client"
	"github.com/carloseduribeiro/Client-Server-API/client/repository/fs"
	"log"
	"os"
)

func main() {
	fsRepository, err := fs.NewTextFileRepository()
	if err != nil {
		log.Fatal(err)
	}
	e, err := client.GetExchange(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	if err = fsRepository.SaveExchange(e); err != nil {
		log.Fatal(err)
	}
	encoder := json.NewEncoder(os.Stdout)
	if err = encoder.Encode(e); err != nil {
		panic(err)
	}
}
