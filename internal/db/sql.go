package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Ensure you import the MySQL driver
)

// InitDB initializes and returns a connection to the database.
func InitDB() *sql.DB {
	// Initialization and connection to the MySQL database.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/groupie")
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	// Check that the database connection is successful.
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	return db
}
