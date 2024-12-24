package middleware

import (
	"net/http"
	"strings"
)

// AuthMiddleware checks if a request contains a valid token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		// Example check for a valid token
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
