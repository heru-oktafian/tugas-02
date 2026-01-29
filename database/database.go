package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
	// Log connection string untuk debugging
	log.Println("Attempting to connect to database...")

	// Open database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return nil, err
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		log.Printf("Connection error: %v", err)
		return nil, err
	}

	// Set connection pool settings (optional tapi recommended)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully")
	return db, nil
}
