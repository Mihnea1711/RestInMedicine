package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
)

// SanitizedResponseWriterContextKey is the context key for the sanitizedResponseWriter.
type SanitizedResponseWriterContextKey int

// Context key for the sanitizedResponseWriter.
const (
	SanitizedResponseWriterContextKeyInstance SanitizedResponseWriterContextKey = iota
)

// SanitizeInputMiddleware is a middleware that sanitizes user input.
func SanitizeInputMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a policy for HTML sanitization
		policy := bluemonday.UGCPolicy()

		// Create a wrapped response writer to intercept and sanitize the response
		sanitizedWriter := newSanitizedResponseWriter(w, policy)

		// Set the sanitized response writer directly in the request context
		ctx := context.WithValue(r.Context(), SanitizedResponseWriterContextKeyInstance, sanitizedWriter)
		r = r.WithContext(ctx)

		log.Println("[RABBIT_SANITIZER] Request URI:", r.RequestURI)
		log.Println("[RABBIT_SANITIZER] Starting input sanitization...")

		// Call the next handler in the chain
		next.ServeHTTP(sanitizedWriter, r)

		log.Println("[RABBIT_SANITIZER] Input sanitization completed.")
	})
}

// SanitizedResponseWriter is a wrapper around http.ResponseWriter that sanitizes the response body.
type SanitizedResponseWriter struct {
	http.ResponseWriter
	policy *bluemonday.Policy
}

// NewSanitizedResponseWriter creates a new SanitizedResponseWriter.
func newSanitizedResponseWriter(w http.ResponseWriter, policy *bluemonday.Policy) *SanitizedResponseWriter {
	return &SanitizedResponseWriter{
		ResponseWriter: w,
		policy:         policy,
	}
}

// Write method intercepts the response body and sanitizes it.
func (sw *SanitizedResponseWriter) Write(b []byte) (int, error) {
	// Sanitize the response body using the policy
	sanitizedBody := sw.policy.Sanitize(string(b))

	log.Println("[RABBIT_SANITIZER] Sanitizing response body...")

	// Write the sanitized body to the original ResponseWriter
	return sw.ResponseWriter.Write([]byte(sanitizedBody))
}
