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
	defer database.DB.Close()

	// Setup router
	r := routes.SetupRouter()

	// Add CORS middleware
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Adjust origin as needed
		handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS"}), // Allowed methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),   // Allowed headers
	)

	// Start the server
	log.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", corsOptions(r)))
}
