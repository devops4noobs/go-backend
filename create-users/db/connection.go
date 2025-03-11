// create-users/db/connection.go
package db

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
		log.Fatal("❌ Failed to connect to database:", err)
		return nil, err
	}

	// Verify database connection
	if err := db.Ping(); err != nil {
		log.Fatal("❌ Database is not reachable:", err)
		return nil, err
	}

	log.Println("✅ Successfully connected to the database!")

	// Auto-create users table
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("❌ Error creating users table:", err)
		return nil, err
	}

	log.Println("✅ Users table is ready!")

	return db, nil
}
