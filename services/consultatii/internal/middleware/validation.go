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

func ValidateConsultationInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var consultation models.Consultation

		contentTypeFlag := isContentTypeJSON(r)
		if !contentTypeFlag {
			errMsg := "Unsupported media type. Content-Type must be application/json"
			log.Printf("[CONSULTATION_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnsupportedMediaType, models.ResponseData{Error: errMsg, Message: "Patient validation failed due to unsupported media type"})
			return
		}

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&consultation)
		decodeFlag, decodeStatus := checkErrorOnDecode(err, w)
		if decodeFlag || decodeStatus != http.StatusOK {
			errMsg := "Failed to decode consultation"
			log.Printf("[CONSULTATION_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Consultation validation failed due to decoding."})
			return
		}

		// Validate the Consultation
		if err := validateConsultation(&consultation); err != nil {
			log.Printf("[CONSULTATION_VALIDATION] Validation error: %s", err)
			utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: err.Error()})
			return
		}

		// If all validations pass, proceed to the actual controller
		ctx := context.WithValue(r.Context(), utils.DECODED_CONSULTATION, &consultation)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Function to check the error when decoding an object
func checkErrorOnDecode(err error, w http.ResponseWriter) (bool, int) {
	if err == nil {
		return false, http.StatusOK
	}
	var errMsg string
	var statusCode int

	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError
	switch {
	// Catch any syntax errors in the JSON and send an error message
	// which interpolates the location of the problem to make it
	// easier for the client to fix.
	case errors.As(err, &syntaxError):
		errMsg = fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
		statusCode = http.StatusUnprocessableEntity

	// In some circumstances Decode() may also return an
	// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
	// is an open issue regarding this at
	// https://github.com/golang/go/issues/25956.
	case errors.Is(err, io.ErrUnexpectedEOF):
		errMsg = "Request body contains badly-formed JSON"
		statusCode = http.StatusUnprocessableEntity

	// Catch any type errors, like trying to assign a string in the
	// JSON request body to a int field in our Person struct. We can
	// interpolate the relevant field name and position into the error
	// message to make it easier for the client to fix.
	case errors.As(err, &unmarshalTypeError):
		errMsg = fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
		statusCode = http.StatusUnprocessableEntity

	// Catch the error caused by extra unexpected fields in the request
	// body. We extract the field name from the error message and
	// interpolate it in our custom error message. There is an open
	// issue regarding turning this into a sentinel error.
	case strings.HasPrefix(err.Error(), "json: unknown field "):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		errMsg = fmt.Sprintf("Request body contains unknown field %s", fieldName)
		statusCode = http.StatusUnprocessableEntity

	// An io.EOF error is returned by Decode() if the request body is
	// empty.
	case errors.Is(err, io.EOF):
		errMsg = "Request body must not be empty"
		statusCode = http.StatusBadRequest

	// Catch the error caused by the request body being too large. Again
	// there is an open issue regarding turning this into a sentinel
	// error at https://github.com/golang/go/issues/30715.
	case err.Error() == "http: request body too large":
		errMsg = "Request body must not be larger than 1MB"
		statusCode = http.StatusRequestEntityTooLarge

	// Otherwise default to logging the error and sending a 500 Internal
	// Server Error response.
	default:
		errMsg = err.Error()
		statusCode = http.StatusInternalServerError
	}

	log.Printf("[PATIENT_VALIDATION] %s", errMsg)
	return true, statusCode
}

func isContentTypeJSON(r *http.Request) bool {
	contentType := r.Header.Get("Content-Type")
	return contentType == "application/json"
}

func validateConsultation(consultation *models.Consultation) error {
	// Validate the IDPatient (assuming it should be greater than 0)
	if consultation.IDPatient <= 0 {
		log.Println("[CONSULTATION_VALIDATION] Invalid patient ID")
		return errors.New("invalid patient ID")
	}

	// Validate the IDDoctor (assuming it should be greater than 0)
	if consultation.IDDoctor <= 0 {
		log.Println("[CONSULTATION_VALIDATION] Invalid doctor ID")
		return errors.New("invalid doctor ID")
	}

	// Validate the Date (assuming it should not be zero)
	if consultation.Date.IsZero() {
		log.Println("[CONSULTATION_VALIDATION] Invalid date")
		return errors.New("invalid date")
	}

	// Validate that the date is not in the past
	currentMoment := time.Now()
	if consultation.Date.Before(currentMoment) {
		log.Println("[CONSULTATION_VALIDATION] Date should not be in the past")
		return errors.New("date should not be in the past")
	}

	// Validate the Diagnostic field (assuming it should not be empty)
	if consultation.Diagnostic == "" {
		log.Println("[CONSULTATION_VALIDATION] Invalid diagnostic")
		return errors.New("invalid diagnostic")
	}

	// Validate each investigation in the list
	for _, inv := range consultation.Investigations {
		if inv.Name == "" {
			log.Println("[CONSULTATION_VALIDATION] Invalid investigation denumire")
			return errors.New("invalid investigation denumire")
		}

		if inv.ProcessingTime <= 0 {
			log.Println("[CONSULTATION_VALIDATION] Invalid investigation durata de procesare")
			return errors.New("invalid investigation durata de procesare")
		}

		if inv.Result == "" {
			log.Println("[CONSULTATION_VALIDATION] Invalid investigation rezultat")
			return errors.New("invalid investigation rezultat")
		}
	}

	// If all validations pass, return nil
	return nil
}
