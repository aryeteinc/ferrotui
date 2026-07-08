package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "ferrotui.db")
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile("internal/db/migrations/001_init.sql")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(string(data))
	if err != nil {
		return nil, err
	}

	return db, nil
}
