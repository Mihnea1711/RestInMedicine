package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		startTime := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// Call the next handler, which can be another middleware or the final handler.
		next.ServeHTTP(w, r)

		log.Printf("Completed %s in %v", r.URL.Path, time.Since(startTime))
	})
}
