package database

import (
	"context"
	"database/sql"
	"github.com/carloseduribeiro/Client-Server-API/server/client"
)

const insertStmt = "INSERT INTO exchange(code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"

type ExchangeRepository struct {
	conn *sql.DB
}

func NewExchangeRepository(conn *sql.DB) *ExchangeRepository {
	return &ExchangeRepository{conn: conn}
}

func (e *ExchangeRepository) Save(ctx context.Context, dto *client.ExchangeDto) error {
	stmt, err := e.conn.PrepareContext(ctx, insertStmt)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, dto.Code, dto.Codein, dto.Name, dto.High, dto.Low, dto.VarBid, dto.PctChange, dto.Bid, dto.Ask, dto.Timestamp, dto.CreateDate)
	if err != nil {
		return err
	}
	return nil
}
