package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)
type Food struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Calories int    `json:"calories"`
}

var db *sql.DB

func loadEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func getEnv(key string) string {
	return os.Getenv(key)
}


func main() {
    // Load environment variables
    loadEnv()
    password := getEnv("PASSWORD")
    connectionString := fmt.Sprintf("root:%s@tcp(localhost:3306)/foodSearch", password)

    // Open the connection to the database
    var err error
    db, err = sql.Open("mysql", connectionString)
    if err != nil {
        log.Fatalf("Error validating sql.Open arguments: %v", err)
    }

    // Verify connection
    if err := db.Ping(); err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    // Set up HTTP routes
    http.HandleFunc("/api/foods", getFoods)

    // Start the HTTP server with CORS enabled
    fmt.Println("Server running on http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:3000"}),
        handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
    )(http.DefaultServeMux)))
}


// getFoods handles GET requests to /api/foods
func getFoods(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request on /api/foods") // Debug log for requests

    rows, err := db.Query("SELECT id, name, calories FROM food")
    if err != nil {
        http.Error(w, "Failed to fetch data from database", http.StatusInternalServerError)
        log.Printf("Error fetching data: %v", err) // Debug log for errors
        return
    }
    defer rows.Close()

    var foods []Food
    for rows.Next() {
        var food Food
        if err := rows.Scan(&food.ID, &food.Name, &food.Calories); err != nil {
            http.Error(w, "Failed to scan data", http.StatusInternalServerError)
            log.Printf("Error scanning row: %v", err) // Debug log for errors
            return
        }
        foods = append(foods, food)
    }

    log.Printf("Fetched data: %v", foods) // Debug log for fetched data

    // Send JSON response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(foods)
}
