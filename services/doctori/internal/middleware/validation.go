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

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&doctor)
		if checkErrorOnDecode(err, w) {
			errMsg := "Failed to decode doctor"
			log.Printf("[DOCTOR_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Doctor validation failed due to decoding."})
			return
		}

		// Basic validation for each field
		if doctor.FirstName == "" || len(doctor.FirstName) > 255 {
			errMsg := "Invalid or missing Nume"
			log.Printf("[DOCTOR_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Doctor validation failed due to first name"})
			return
		}

		if doctor.SecondName == "" || len(doctor.SecondName) > 255 {
			errMsg := "Invalid or missing Prenume"
			log.Printf("[DOCTOR_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Doctor validation failed due to second name"})
			return
		}

		// Validate email format using regex
		if !utils.EmailRegex.MatchString(doctor.Email) || len(doctor.Email) > 255 {
			errMsg := "Invalid or missing Email"
			log.Printf("[DOCTOR_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Doctor validation failed due to email"})
			return
		}

		// Validate Romanian phone number format
		if !utils.PhoneRegex.MatchString(doctor.PhoneNumber) || len(doctor.PhoneNumber) != 10 {
			errMsg := "Invalid or missing Telefon"
			log.Printf("[DOCTOR_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Doctor validation failed due to phone nr"})
			return
		}

		if !isValidSpecializare(models.Specialization(doctor.Specialization)) {
			errMsg := "Invalid Specializare in request"
			log.Printf("[DOCTOR_VALIDATION] %s: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Doctor validation failed due to specialization"})
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

	log.Printf("[DOCTOR_VALIDATION] %s", errMsg)
	http.Error(w, errMsg, http.StatusBadRequest)
	return true
}

func ValidateEmail(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		email, ok := vars[utils.FETCH_DOCTOR_BY_EMAIL_PARAMETER]
		if !ok {
			errMsg := "Email not provided in request"
			log.Printf("[DOCTOR_VALIDATION] Email not provided in request: %s", r.RequestURI)
			utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg})
			return
		}

		if !utils.EmailRegex.MatchString(email) {
			errMsg := "Invalid email format"
			log.Printf("[DOCTOR_VALIDATION] %s for email: %s in request: %s", errMsg, email, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg})
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
		if checkErrorOnDecode(err, w) {
			errMsg := "Failed to decode doctor"
			log.Printf("[DOCTOR_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Doctor validation failed due to decoding."})
			return
		}

		// Basic validation for each field
		if !isBoolean(doctorActivityData.IsActive) {
			errMsg := "Invalid IsActive field type"
			log.Printf("[DOCTOR_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Doctor validation failed due to IsActive"})
			return
		}

		if doctorActivityData.IDUser <= 0 {
			errMsg := "Invalid or missing IDUser"
			log.Printf("[DOCTOR_VALIDATION] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Doctor validation failed due to first name"})
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
