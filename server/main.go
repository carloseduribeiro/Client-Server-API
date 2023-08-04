package main

import (
	"github.com/carloseduribeiro/Client-Server-API/server/api"
	"github.com/carloseduribeiro/Client-Server-API/server/client/exchange"
	"github.com/carloseduribeiro/Client-Server-API/server/database"
	"net/http"
)

const (
	dbFileName = "database.db"
)

func main() {
	conn, err := database.NewSQLiteConnection(dbFileName)
	if err != nil {
		panic(err)
	}
	exchangeRepository := database.NewExchangeRepository(conn)
	handler := api.NewExchangeHandler(exchange.GetExchange, exchangeRepository)
	http.HandleFunc("/cotacao", handler.GetExchange)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
