// create-users/db/connection.go
package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
		return nil, err
	}

	// Verify database connection
	if err := db.Ping(); err != nil {
		log.Fatal("❌ Database is not reachable:", err)
		return nil, err
	}

	log.Println("✅ Successfully connected to the database!")

	return db, nil
}
