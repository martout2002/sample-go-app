package main

import (
	"log"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/routes"
	"github.com/gorilla/handlers"
)

func main() {
	// Initialize the database
	log.Println("Initializing database connection")
	database.InitDB()
	defer database.DB.Close() // Ensure the database connection is closed when the app shuts down

	// Set up the router
	router := routes.SetupRouter()

	// Enable CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	// Start the HTTP server
	log.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", corsHandler(router)))
}
