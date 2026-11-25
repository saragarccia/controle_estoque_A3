package config

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB() (*sql.DB, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, errors.New("variável de ambiente DATABASE_URL não definida")
	}

	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
