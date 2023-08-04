package database

import (
	"context"
	"database/sql"
	"github.com/carloseduribeiro/Client-Server-API/server/client/exchange"
	"time"
)

const (
	databaseTimeout = 10 * time.Millisecond
	insertStmt      = "INSERT INTO exchange(code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
)

type ExchangeRepository struct {
	conn *sql.DB
}

func NewExchangeRepository(conn *sql.DB) *ExchangeRepository {
	return &ExchangeRepository{conn: conn}
}

func (e *ExchangeRepository) Save(ctx context.Context, dto *exchange.Exchange) error {
	persistCtx, cancel := context.WithTimeout(ctx, databaseTimeout)
	defer cancel()
	stmt, err := e.conn.PrepareContext(persistCtx, insertStmt)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(persistCtx, dto.Code, dto.Codein, dto.Name, dto.High, dto.Low, dto.VarBid, dto.PctChange, dto.Bid, dto.Ask, dto.Timestamp, dto.CreateDate)
	if err != nil {
		return err
	}
	return nil
}
