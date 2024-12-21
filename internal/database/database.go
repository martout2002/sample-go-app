package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB is the global database connection
var DB *sql.DB

// InitDB initializes the database connection
func InitDB() {
	// Load environment variables
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	password := os.Getenv("PASSWORD")
	connectionString := fmt.Sprintf("root:%s@tcp(localhost:3306)/foodSearch", password)

	var err error
	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Verify connection
	if err := DB.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Database connection established")
}

// GetDB returns the current database connection
func GetDB() (*sql.DB, error) {
	if DB == nil {
		return nil, fmt.Errorf("database connection not initialized")
	}
	return DB, nil
}