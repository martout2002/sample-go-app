package routes

import (
	"net/http"

	"github.com/CVWO/sample-go-app/internal/controllers"
	"github.com/CVWO/sample-go-app/internal/middleware"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Threads
	mux.Handle("/api/threads", middleware.LoggingMiddleware(http.HandlerFunc(controllers.GetThreads)))
	mux.Handle("/api/threads/add", middleware.LoggingMiddleware(http.HandlerFunc(controllers.CreateThread)))
	mux.Handle("/api/threads/delete", middleware.LoggingMiddleware(http.HandlerFunc(controllers.DeleteThread)))
	mux.Handle("/api/threads/like", middleware.LoggingMiddleware(http.HandlerFunc(controllers.LikeThread)))
	mux.Handle("/api/threads/likes", middleware.LoggingMiddleware(http.HandlerFunc(controllers.GetLikesCount)))

	// Users
	mux.Handle("/api/users", middleware.LoggingMiddleware(http.HandlerFunc(controllers.HandleListUsers)))
	mux.Handle("/api/signup", middleware.LoggingMiddleware(http.HandlerFunc(controllers.HandleSignup)))
	mux.Handle("/api/login", middleware.LoggingMiddleware(http.HandlerFunc(controllers.HandleLogin)))

	return mux
}
