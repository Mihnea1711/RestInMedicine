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

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func ValidateDoctorInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var doctor models.Doctor

		contentTypeFlag := isContentTypeJSON(r)
		if !contentTypeFlag {
			errMsg := "Unsupported media type. Content-Type must be application/json"
			log.Printf("[DOCTOR_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnsupportedMediaType, models.ResponseData{Error: errMsg, Message: "Patient validation failed due to unsupported media type"})
			return
		}

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&doctor)
		decodeFlag, decodeStatus := checkErrorOnDecode(err, w)
		if decodeFlag || decodeStatus != http.StatusOK {
			errMsg := "Failed to decode doctor"
			log.Printf("[DOCTOR_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, decodeStatus, models.ResponseData{Error: errMsg, Message: "Doctor validation failed due to decoding."})
			return
		}

		// Basic validation for each field
		if doctor.FirstName == "" || len(doctor.FirstName) > utils.MaxNameLength {
			errMsg := "Invalid or missing Nume"
			log.Printf("[DOCTOR_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{Error: errMsg, Message: "Doctor validation failed due to first name"})
			return
		}

		if doctor.SecondName == "" || len(doctor.SecondName) > utils.MaxNameLength {
			errMsg := "Invalid or missing Prenume"
			log.Printf("[DOCTOR_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{Error: errMsg, Message: "Doctor validation failed due to second name"})
			return
		}

		// Validate email format using regex
		if !utils.EmailRegex.MatchString(doctor.Email) || len(doctor.Email) > utils.MaxEmailLength {
			errMsg := "Invalid or missing Email"
			log.Printf("[DOCTOR_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{Error: errMsg, Message: "Doctor validation failed due to email"})
			return
		}

		// Validate Romanian phone number format
		if !utils.PhoneRegex.MatchString(doctor.PhoneNumber) || len(doctor.PhoneNumber) != utils.PhoneNumberLength {
			errMsg := "Invalid or missing Telefon"
			log.Printf("[DOCTOR_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{Error: errMsg, Message: "Doctor validation failed due to phone nr"})
			return
		}

		if !isValidSpecializare(models.Specialization(doctor.Specialization)) {
			errMsg := "Invalid Specializare in request"
			log.Printf("[DOCTOR_VALIDATION] %s: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{Error: errMsg, Message: "Doctor validation failed due to specialization"})
			return
		}

		log.Printf("[DOCTOR_VALIDATION] Doctor info validated successfully in request: %s", r.RequestURI)

		// If all validations pass, proceed to the actual controller
		ctx := context.WithValue(r.Context(), utils.DECODED_DOCTOR, &doctor)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func isValidSpecializare(specializare models.Specialization) bool {
	log.Printf("[DOCTOR_VALIDATION] Checking validity of specializare: %v...", specializare)
	for _, validSpec := range utils.ValidSpecializations {
		if specializare == validSpec {
			return true
		}
	}
	log.Printf("[DOCTOR_VALIDATION] Specializare %v is not valid.", specializare)
	return false
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

func ValidateEmail(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		email, ok := vars[utils.FETCH_DOCTOR_BY_EMAIL_PARAMETER]
		if !ok {
			errMsg := "Email not provided in request"
			log.Printf("[DOCTOR_VALIDATION] Email not provided in request: %s", r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{Error: errMsg})
			return
		}

		if !utils.EmailRegex.MatchString(email) || len(email) > utils.MaxEmailLength {
			errMsg := "Invalid email format"
			log.Printf("[DOCTOR_VALIDATION] %s for email: %s in request: %s", errMsg, email, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{Error: errMsg})
			return
		}

		log.Printf("[DOCTOR_VALIDATION] Email validated successfully: %s", email)
		next.ServeHTTP(w, r)
	})
}

func ValidateDoctorActivityInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var doctorActivityData models.ActivityData

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&doctorActivityData)
		decodeFlag, decodeStatus := checkErrorOnDecode(err, w)
		if decodeFlag || decodeStatus != http.StatusOK {
			errMsg := "Failed to decode doctor"
			log.Printf("[DOCTOR_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, decodeStatus, models.ResponseData{Error: errMsg, Message: "Doctor validation failed due to decoding."})
			return
		}

		// Basic validation for each field
		if !isBoolean(doctorActivityData.IsActive) {
			errMsg := "Invalid IsActive field type"
			log.Printf("[DOCTOR_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{Error: errMsg, Message: "Doctor validation failed due to IsActive"})
			return
		}

		log.Printf("[DOCTOR_VALIDATION] Doctor info validated successfully in request: %s", r.RequestURI)

		// If all validations pass, proceed to the actual controller
		ctx := context.WithValue(r.Context(), utils.DECODED_DOCTOR_ACTIVITY, &doctorActivityData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// isBoolean checks if a variable is of boolean type
func isBoolean(value interface{}) bool {
	return reflect.ValueOf(value).Kind() == reflect.Bool
}
