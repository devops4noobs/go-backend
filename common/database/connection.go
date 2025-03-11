// common/database/connection.go
package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	//dbURL := os.Getenv("DATABASE_URL")
	const dbURL = "postgres://postgres:postgres@postgres.czw08msma20l.eu-central-1.rds.amazonaws.com:5432/postgres"
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
