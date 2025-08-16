package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewConnection(connectionString string) (*sql.DB, error) {
	if connectionString == "" {
		connectionString = ":memory:"
	}

	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT
    )`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
