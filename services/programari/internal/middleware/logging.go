package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger is a middleware that logs the start and end of each request, along with some useful data about the request.
func RouteLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Continue to the next handler.
		next.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s\n",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}
