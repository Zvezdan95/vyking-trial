package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// ConnectDB establishes a connection to the MySQL database
func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	dbUser := "root"
	dbPass := os.Getenv("DB_ROOT_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := "db"

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbName)

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}
	// Wait for MySQL to start (new addition)
	for i := 0; i < 30; i++ { // Try for 30 seconds
		err = db.Ping()
		if err == nil {
			break // MySQL is ready
		}
		log.Println("Waiting for MySQL to start...")
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return db, nil
}
