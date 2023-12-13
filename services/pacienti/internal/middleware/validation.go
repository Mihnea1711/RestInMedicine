package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func ValidatePacientInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var patient models.Patient

		contentTypeFlag := isContentTypeJSON(r)
		if !contentTypeFlag {
			errMsg := "Unsupported media type. Content-Type must be application/json"
			log.Printf("[PATIENT_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnsupportedMediaType, models.ResponseData{Error: errMsg, Message: "Patient validation failed due to unsupported media type"})
			return
		}

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&patient)
		decodeFlag, decodeStatus := checkErrorOnDecode(err, w)
		if decodeFlag || decodeStatus != http.StatusOK {
			errMsg := "Failed to decode patient"
			log.Printf("[PATIENT_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, decodeStatus, models.ResponseData{Error: errMsg, Message: "Patient validation failed due to decoding."})
			return
		}

		// Basic validation for each field
		if patient.FirstName == "" || len(patient.FirstName) > utils.MaxNameLength {
			errMsg := "Invalid or missing Nume"
			log.Printf("[PATIENT_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{Error: errMsg, Message: "Patient validation failed due to first name"})
			return
		}

		if patient.SecondName == "" || len(patient.SecondName) > utils.MaxNameLength {
			errMsg := "Invalid or missing Prenume"
			log.Printf("[PATIENT_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{Error: errMsg, Message: "Patient validation failed due to second name"})
			return
		}

		// Validate email format using regex
		if !utils.EmailRegex.MatchString(patient.Email) || len(patient.Email) > utils.MaxEmailLength {
			errMsg := "Invalid or missing Email"
			log.Printf("[PATIENT_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{Error: errMsg, Message: "Patient validation failed due to email"})
			return
		}

		// Validate Romanian phone number format
		if !utils.PhoneRegex.MatchString(patient.PhoneNumber) || len(patient.PhoneNumber) != utils.PhoneNumberLength {
			errMsg := "Invalid or missing Telefon"
			log.Printf("[PATIENT_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{Error: errMsg, Message: "Patient validation failed due to phone nr"})
			return
		}

		// Check if DataNasterii is valid (18 years in the past)
		minimumBirthDate := time.Now().AddDate(-18, 0, 0)
		if patient.BirthDay.After(minimumBirthDate) {
			errMsg := "Invalid birthday (must be at least 18 years ago)"
			log.Printf("[PATIENT_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{Error: errMsg, Message: "Patient validation failed due to birthdate"})
			return
		}

		// In your ValidatePacientInfo middleware
		if ok, errMsg := validateCNPBirthdate(patient.CNP, patient.BirthDay); !ok {
			log.Printf("[PATIENT_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{Error: errMsg, Message: "Patient validation failed due to cnp"})
			return
		}

		log.Printf("[PATIENT_VALIDATION] Patient info validated successfully in request: %s", r.RequestURI)

		// If all validations pass, proceed to the actual controller
		ctx := context.WithValue(r.Context(), utils.DECODED_PATIENT, &patient)
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

// Function to validate the CNP birthdate against DataNasterii
func validateCNPBirthdate(cnp string, dataNasterii time.Time) (bool, string) {
	if len(cnp) != utils.MaxCNPLength || cnp == "" {
		errMsg := "Incorrect CNP format"
		log.Printf("[PATIENT_VALIDATION] %s", errMsg)
		return false, errMsg
	}

	// Extract birthdate from CNP
	cnpBirthdateStr := cnp[1:7] // Extract the 6-digit date of birth from the CNP
	cnpBirthdate, err := time.Parse(utils.CNP_DATE_FORMAT, cnpBirthdateStr)
	if err != nil {
		errMsg := "CNP not matching birthdate"
		log.Printf("[PATIENT_VALIDATION] %s", errMsg)
		return false, errMsg
	}

	// Compare the extracted CNP birthdate with DataNasterii
	if !cnpBirthdate.Equal(dataNasterii) {
		errMsg := "CNP birthdate does not match DataNasterii"
		log.Printf("[PATIENT_VALIDATION] %s", errMsg)
		return false, errMsg
	}

	return true, ""
}

// Middleware to validate the email param in the request
func ValidateEmail(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		email, ok := vars[utils.FETCH_PATIENT_BY_EMAIL_PARAMETER]
		if !ok {
			errMsg := "Email not provided in request"
			log.Printf("[PATIENT_VALIDATION] %s: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{Error: errMsg})
			return
		}

		if !utils.EmailRegex.MatchString(email) {
			errMsg := "Invalid email format"
			log.Printf("[PATIENT_VALIDATION] %s for email: %s in request: %s", errMsg, email, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{Error: errMsg})
			return
		}

		log.Printf("[PATIENT_VALIDATION] Email validated successfully: %s", email)
		next.ServeHTTP(w, r)
	})
}

func ValidatePatientActivityInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var PatientActivityData models.ActivityData

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&PatientActivityData)
		decodeFlag, decodeStatus := checkErrorOnDecode(err, w)
		if decodeFlag || decodeStatus != http.StatusOK {
			errMsg := "Failed to decode Patient"
			log.Printf("[PATIENT_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, decodeStatus, models.ResponseData{Error: errMsg, Message: "Patient validation failed due to decoding."})
			return
		}

		// Basic validation for each field
		if !isBoolean(PatientActivityData.IsActive) {
			errMsg := "Invalid IsActive field type"
			log.Printf("[PATIENT_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{Error: errMsg, Message: "Patient validation failed due to IsActive"})
			return
		}

		log.Printf("[PATIENT_VALIDATION] Patient info validated successfully in request: %s", r.RequestURI)

		// If all validations pass, proceed to the actual controller
		ctx := context.WithValue(r.Context(), utils.DECODED_PATIENT_ACTIVITY, &PatientActivityData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// isBoolean checks if a variable is of boolean type
func isBoolean(value interface{}) bool {
	return reflect.ValueOf(value).Kind() == reflect.Bool
}
