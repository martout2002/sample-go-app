package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Thread represents a food review thread
type Thread struct {
	Title         string
	StoreName     string
	StoreLocation string
	AuthorName    string
	Details       string
	Rating        float64
	Comments      string
}

// Load environment variables
func loadEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func getEnv(key string) string {
	return os.Getenv(key)
}

func main() {
	loadEnv()
	password := getEnv("PASSWORD")
	connectionString := fmt.Sprintf("root:%s@tcp(localhost:3306)/foodSearch", password)

	// Connect to the database
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Insert hardcoded food review threads
	insertFoodReviewThreads(db)
}

func insertFoodReviewThreads(db *sql.DB) {
	threads := []Thread{
		{
			Title:         "Amazing Pizza at Mario's",
			StoreName:     "Mario's Pizzeria",
			StoreLocation: "123 Main Street, New York",
			AuthorName:    "John Doe",
			Details:       "The pepperoni pizza was out of this world! Crispy crust and fresh toppings.",
			Rating:        4.8,
			Comments:      "Highly recommend for pizza lovers!",
		},
		{
			Title:         "Great Coffee and Ambience",
			StoreName:     "Star Coffee",
			StoreLocation: "456 Elm Street, San Francisco",
			AuthorName:    "Jane Smith",
			Details:       "The cappuccino was perfect, and the quiet environment was ideal for working.",
			Rating:        4.5,
			Comments:      "Nice spot for a peaceful coffee break.",
		},
		{
			Title:         "Sushi Heaven at Tokyo Bites",
			StoreName:     "Tokyo Bites",
			StoreLocation: "789 Sakura Avenue, Tokyo",
			AuthorName:    "Alex Tan",
			Details:       "The sushi rolls were fresh and beautifully presented. The salmon nigiri was a highlight.",
			Rating:        5.0,
			Comments:      "Best sushi I've had outside of Japan!",
		},
	}

	for _, thread := range threads {
		query := `INSERT INTO threads (title, store_name, store_location, author_name, details, rating, comments) 
                  VALUES (?, ?, ?, ?, ?, ?, ?)`
		_, err := db.Exec(query, thread.Title, thread.StoreName, thread.StoreLocation, thread.AuthorName, thread.Details, thread.Rating, thread.Comments)
		if err != nil {
			log.Printf("Failed to insert thread %s: %v", thread.Title, err)
			continue
		}
		log.Printf("Inserted: %s | %s | %s", thread.Title, thread.StoreName, thread.StoreLocation)
	}

	log.Println("Inserted 3 food review threads successfully.")
}
