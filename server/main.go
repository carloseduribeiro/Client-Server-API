package main

import (
	"github.com/carloseduribeiro/Client-Server-API/server/api"
	"github.com/carloseduribeiro/Client-Server-API/server/client"
	"github.com/carloseduribeiro/Client-Server-API/server/database"
	"net/http"
)

const (
	dbFileName      = "database.db"
	exchangeAPIHost = "https://economia.awesomeapi.com.br"
)

func main() {
	conn, err := database.NewSQLiteConnection(dbFileName)
	if err != nil {
		panic(err)
	}
	exchangeClient := client.NewExchangeClient(exchangeAPIHost, http.DefaultClient)
	exchangeRepository := database.NewExchangeRepository(conn)
	handler := api.NewExchangeHandler(exchangeClient, exchangeRepository)
	http.HandleFunc("/cotacao", handler.GetExchange)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
