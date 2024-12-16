package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/routes"

	"github.com/gorilla/handlers"
)

func main() {
	// Initialize the database
	database.InitDB()
	defer database.DB.Close()

	// Setup routes
	router := routes.SetupRouter()

	// Start the HTTP server
	fmt.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	)(router)))
}
