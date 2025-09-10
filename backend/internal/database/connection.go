package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func NewConnection(connectionString string) *Database {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Successfully connected to database")
	return &Database{DB: db}
}

func (d *Database) Close() error {
	return d.DB.Close()
}
