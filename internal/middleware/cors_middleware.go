package middleware

import (
	"log"
	"net/http"
)

// CORS adds the appropriate headers for cross-origin requests.
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Username")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			log.Printf("CORS Preflight Request: %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass to the next handler
		next.ServeHTTP(w, r)
	})
}