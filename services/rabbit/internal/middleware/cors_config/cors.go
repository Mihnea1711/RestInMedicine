package cors_config

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func SetupCORS(router *mux.Router) {
	// CORS middleware with custom options
	corsOptions := cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}

	corsMiddleware := cors.New(corsOptions)

	// Handle OPTIONS requests for CORS pre-flight
	router.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusOK)
	})

	// Use CORS middleware
	router.Use(corsMiddleware.Handler)
}
