package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func ValidatePacientInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var pacient models.Pacient

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&pacient)
		if checkErrorOnDecode(err, w) {
			log.Printf("[MIDDLEWARE] Failed to decode pacient in request: %s", r.RequestURI)
			http.Error(w, "Failed to decode pacient", http.StatusBadRequest)
			return
		}

		// Basic validation for each field
		if pacient.Nume == "" || len(pacient.Nume) > 255 {
			log.Printf("[MIDDLEWARE] Invalid or missing Nume in request: %s", r.RequestURI)
			http.Error(w, "Invalid or missing Nume", http.StatusBadRequest)
			return
		}

		if pacient.Prenume == "" || len(pacient.Prenume) > 255 {
			log.Printf("[MIDDLEWARE] Invalid or missing Prenume in request: %s", r.RequestURI)
			http.Error(w, "Invalid or missing Prenume", http.StatusBadRequest)
			return
		}

		// Validate email format using regex
		if !utils.EmailRegex.MatchString(pacient.Email) || len(pacient.Email) > 255 {
			log.Printf("[MIDDLEWARE] Invalid or missing Email in request: %s", r.RequestURI)
			http.Error(w, "Invalid or missing Email", http.StatusBadRequest)
			return
		}

		// Validate Romanian phone number format
		if !utils.PhoneRegex.MatchString(pacient.Telefon) || len(pacient.Telefon) != 10 {
			log.Printf("[MIDDLEWARE] Invalid or missing Telefon in request: %s", r.RequestURI)
			http.Error(w, "Invalid or missing Telefon", http.StatusBadRequest)
			return
		}

		// Check if DataNasterii is valid (18 years in the past)
		minimumBirthDate := time.Now().AddDate(-18, 0, 0)
		if pacient.DataNasterii.After(minimumBirthDate) {
			log.Printf("[MIDDLEWARE] Invalid DataNasterii (must be at least 18 years ago) in request: %s", r.RequestURI)
			http.Error(w, "Invalid DataNasterii (must be at least 18 years ago)", http.StatusBadRequest)
			return
		}

		// In your ValidatePacientInfo middleware
		if !validateCNPBirthdate(pacient.CNP, pacient.DataNasterii) {
			log.Printf("[MIDDLEWARE] CNP birthdate does not match DataNasterii in request: %s", r.RequestURI)
			http.Error(w, "CNP birthdate does not match DataNasterii", http.StatusBadRequest)
			return
		}

		log.Printf("[MIDDLEWARE] Pacient info validated successfully in request: %s", r.RequestURI)

		// If all validations pass, proceed to the actual controller
		ctx := context.WithValue(r.Context(), utils.DECODED_PACIENT, &pacient)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func checkErrorOnDecode(err error, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}
	var errMsg string

	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError
	switch {
	// Catch any syntax errors in the JSON and send an error message
	// which interpolates the location of the problem to make it
	// easier for the client to fix.
	case errors.As(err, &syntaxError):
		errMsg = fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)

	// In some circumstances Decode() may also return an
	// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
	// is an open issue regarding this at
	// https://github.com/golang/go/issues/25956.
	case errors.Is(err, io.ErrUnexpectedEOF):
		errMsg = "Request body contains badly-formed JSON"

	// Catch any type errors, like trying to assign a string in the
	// JSON request body to a int field in our Person struct. We can
	// interpolate the relevant field name and position into the error
	// message to make it easier for the client to fix.
	case errors.As(err, &unmarshalTypeError):
		errMsg = fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)

	// Catch the error caused by extra unexpected fields in the request
	// body. We extract the field name from the error message and
	// interpolate it in our custom error message. There is an open
	// issue regarding turning this into a sentinel error.
	case strings.HasPrefix(err.Error(), "json: unknown field "):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		errMsg = fmt.Sprintf("Request body contains unknown field %s", fieldName)

	// An io.EOF error is returned by Decode() if the request body is
	// empty.
	case errors.Is(err, io.EOF):
		errMsg = "Request body must not be empty"

	// Catch the error caused by the request body being too large. Again
	// there is an open issue regarding turning this into a sentinel
	// error at https://github.com/golang/go/issues/30715.
	case err.Error() == "http: request body too large":
		errMsg = "Request body must not be larger than 1MB"

	// Otherwise default to logging the error and sending a 500 Internal
	// Server Error response.
	default:
		errMsg = err.Error()
	}

	log.Printf("[MIDDLEWARE] %s", errMsg)
	http.Error(w, errMsg, http.StatusBadRequest)
	return true
}

// Function to validate the CNP birthdate against DataNasterii
func validateCNPBirthdate(cnp string, dataNasterii time.Time) bool {
	if len(cnp) != 13 || cnp == "" {
		log.Println("[MIDDLEWARE] Incorrect CNP format")
		return false
	}

	// Extract birthdate from CNP
	cnpBirthdateStr := cnp[1:7] // Extract the 6-digit date of birth from the CNP
	cnpBirthdate, err := time.Parse("060102", cnpBirthdateStr)
	if err != nil {
		log.Println("[MIDDLEWARE] CNP not matching birthdate")
		return false
	}

	// Compare the extracted CNP birthdate with DataNasterii
	return cnpBirthdate.Equal(dataNasterii)
}

func ValidateEmail(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		email, ok := vars["email"]
		if !ok {
			log.Println("[MIDDLEWARE] Email not provided in request:", r.RequestURI)
			http.Error(w, "Email not provided", http.StatusBadRequest)
			return
		}

		if !utils.EmailRegex.MatchString(email) {
			log.Printf("[MIDDLEWARE] Invalid email format for email: %s in request: %s", email, r.RequestURI)
			http.Error(w, "Invalid email format", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// add similar validation middlewares for Update, Delete, etc.
