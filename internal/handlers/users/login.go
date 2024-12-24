package users

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/database"
)

// LoginRequest represents the login request payload
type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// LoginResponse represents the login response payload
type LoginResponse struct {
    Message string `json:"message"`
    Success bool   `json:"success"`
}

// HandleLogin verifies username and password
func HandleLogin(w http.ResponseWriter, r *http.Request) {
    var loginReq LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    var storedPassword string
    err := database.DB.QueryRow("SELECT password FROM users WHERE username = ?", loginReq.Username).Scan(&storedPassword)
    if err == sql.ErrNoRows {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    } else if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    // Compare the provided password directly
    if storedPassword != loginReq.Password {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    // Successful login
    json.NewEncoder(w).Encode(LoginResponse{
        Message: "Login successful",
        Success: true,
    })

	
}
