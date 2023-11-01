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

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func ValidateDoctorInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var doctor models.Doctor

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&doctor)
		if checkErrorOnDecode(err, w) {
			return
		}

		// Basic validation for each field
		if doctor.Nume == "" || len(doctor.Nume) > 255 {
			http.Error(w, "Invalid or missing Nume", http.StatusBadRequest)
			return
		}

		if doctor.Prenume == "" || len(doctor.Prenume) > 255 {
			http.Error(w, "Invalid or missing Prenume", http.StatusBadRequest)
			return
		}

		// Validate email format using regex
		if !utils.EmailRegex.MatchString(doctor.Email) || len(doctor.Email) > 255 {
			http.Error(w, "Invalid or missing Email", http.StatusBadRequest)
			return
		}

		// Validate Romanian phone number format
		if !utils.PhoneRegex.MatchString(doctor.Telefon) || len(doctor.Telefon) != 10 {
			http.Error(w, "Invalid or missing Telefon", http.StatusBadRequest)
			return
		}

		if !isValidSpecializare(models.Specializare(doctor.Specializare)) {
			http.Error(w, "Invalid Specializare", http.StatusBadRequest)
			return
		}

		// If all validations pass, proceed to the actual controller
		ctx := context.WithValue(r.Context(), utils.DECODED_DOCTOR, &doctor)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func isValidSpecializare(specializare models.Specializare) bool {
	for _, validSpec := range utils.ValidSpecializari {
		if specializare == validSpec {
			return true
		}
	}
	return false
}

func checkErrorOnDecode(err error, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError
	switch {
	// Catch any syntax errors in the JSON and send an error message
	// which interpolates the location of the problem to make it
	// easier for the client to fix.
	case errors.As(err, &syntaxError):
		msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)

	// In some circumstances Decode() may also return an
	// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
	// is an open issue regarding this at
	// https://github.com/golang/go/issues/25956.
	case errors.Is(err, io.ErrUnexpectedEOF):
		msg := "Request body contains badly-formed JSON"
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)

	// Catch any type errors, like trying to assign a string in the
	// JSON request body to a int field in our Person struct. We can
	// interpolate the relevant field name and position into the error
	// message to make it easier for the client to fix.
	case errors.As(err, &unmarshalTypeError):
		msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)

	// Catch the error caused by extra unexpected fields in the request
	// body. We extract the field name from the error message and
	// interpolate it in our custom error message. There is an open
	// issue at https://github.com/golang/go/issues/29035 regarding
	// turning this into a sentinel error.
	case strings.HasPrefix(err.Error(), "json: unknown field "):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)

	// An io.EOF error is returned by Decode() if the request body is
	// empty.
	case errors.Is(err, io.EOF):
		msg := "Request body must not be empty"
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)

	// Catch the error caused by the request body being too large. Again
	// there is an open issue regarding turning this into a sentinel
	// error at https://github.com/golang/go/issues/30715.
	case err.Error() == "http: request body too large":
		msg := "Request body must not be larger than 1MB"
		log.Println(msg)
		http.Error(w, msg, http.StatusRequestEntityTooLarge)

	// Otherwise default to logging the error and sending a 500 Internal
	// Server Error response.
	default:
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	return true
}

func ValidateEmail(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		email, ok := vars["email"]
		if !ok {
			log.Println("Email not provided in request:", r.RequestURI)
			http.Error(w, "Email not provided", http.StatusBadRequest)
			return
		}

		if !utils.EmailRegex.MatchString(email) {
			log.Printf("Invalid email format for email: %s in request: %s", email, r.RequestURI)
			http.Error(w, "Invalid email format", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// add similar validation middlewares for Update, Delete, etc.
