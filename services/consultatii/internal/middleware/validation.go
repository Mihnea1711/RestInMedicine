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

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
)

func ValidateConsultatieInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var consultatie models.Consultatie

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&consultatie)
		if checkErrorOnDecode(err, w) {
			return
		}

		// Validate the Consultatie
		if err := validateConsultatie(&consultatie); err != nil {
			// Handle the validation error, e.g., return an error response
			log.Printf("[MIDDLEWARE] Validation error: %s", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// If all validations pass, proceed to the actual controller
		ctx := context.WithValue(r.Context(), utils.DECODED_CONSULTATIE, &consultatie)
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

func validateConsultatie(consultatie *models.Consultatie) error {
	if consultatie.IDPacient <= 0 {
		return errors.New("invalid pacient id")
	}

	if consultatie.IDDoctor <= 0 {
		return errors.New("invalid doctor id")
	}

	if consultatie.Date.IsZero() {
		return errors.New("invalid date")
	}

	currentMoment := time.Now()
	if consultatie.Date.Before(currentMoment) {
		return errors.New("date should not be in the past")
	}

	if consultatie.Diagnostic == "" {
		return errors.New("invalid diagnostic")
	}

	for _, inv := range consultatie.Investigatii {
		if inv.Denumire == "" {
			return errors.New("invalid investigatie denumire")
		}

		if inv.DurataProcesare <= 0 {
			return errors.New("invalid investigatie durata de procesare")
		}

		if inv.Rezultat == "" {
			return errors.New("invalid investigatie rezultat")
		}
	}

	return nil
}
