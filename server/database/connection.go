package database

import (
	"database/sql"
	"embed"
	_ "github.com/mattn/go-sqlite3"
	"io"
)

//go:embed schema.sql
var schema embed.FS

// NewSQLiteConnection creates a sql.DB instance and execute the statements inside schema.sql to create table
func NewSQLiteConnection(fileName string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", fileName)
	if err != nil {
		return nil, err
	}
	schemaFile, err := schema.Open("schema.sql")
	if err != nil {
		return nil, err
	}
	defer schemaFile.Close()
	var statements []byte
	statements, err = io.ReadAll(schemaFile)
	if err != nil {
		return nil, err
	}
	if _, err = conn.Exec(string(statements)); err != nil {
		return nil, err
	}
	return conn, nil
}
