package main

import (
	"log"
	"net/http"

	"searchFoodBackend/internal/database"
	"searchFoodBackend/internal/middleware"
	"searchFoodBackend/internal/routes"
)

func main() {
	// Initialize the database
	log.Println("Initializing database connection")
	database.InitDB()
	defer database.DB.Close()

	// Setup router
	r := routes.SetupRouter()

	// Wrap router with global CORS middleware
	withCORS := middleware.CORS(r)

	// Start the server
	log.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", withCORS))
}
