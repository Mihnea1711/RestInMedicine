package validation

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// ValidatePatientData validates the PatientData struct using the validator package
func validatePatientData(patientData models.PatientData) error {
	validate := validator.New()
	return validate.Struct(patientData)
}

// ValidatePatientData is a middleware that validates PatientData
func ValidatePatientData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var patientData models.PatientData

		contentTypeFlag := isContentTypeJSON(r)
		if !contentTypeFlag {
			errMsg := "Unsupported media type. Content-Type must be application/json"
			log.Printf("[MIDDLEWARE_GATEWAY] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnsupportedMediaType, models.ResponseData{Error: errMsg, Message: "Patient validation failed due to unsupported media type"})
			return
		}

		// Decode the request body into PatientData
		err := json.NewDecoder(r.Body).Decode(&patientData)
		if err != nil {
			logAndRespondWithError(w, http.StatusUnprocessableEntity, "Error decoding patient request body", err)
			return
		}

		// Validate PatientData
		if err := validatePatientData(patientData); err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Validation error for patient struct", err)
			return
		}

		// If validation passes, proceed to the next handler
		ctx := context.WithValue(r.Context(), utils.DECODED_PATIENT_DATA, &patientData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ValidatePatientData validates the PatientData struct using the validator package
func validatePatientActivityData(patientActivityData models.ActivityData) error {
	validate := validator.New()
	return validate.Struct(patientActivityData)
}

// ValidatePatientData is a middleware that validates PatientData
func ValidatePatientActivityData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var patientActivityData models.ActivityData

		log.Println(r.Body)

		// Decode the request body into PatientData
		err := json.NewDecoder(r.Body).Decode(&patientActivityData)
		if err != nil {
			logAndRespondWithError(w, http.StatusUnprocessableEntity, "Error decoding patient request body", err)
			return
		}

		// Validate PatientData
		if err := validatePatientActivityData(patientActivityData); err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Validation error for patient struct", err)
			return
		}

		// If validation passes, proceed to the next handler
		ctx := context.WithValue(r.Context(), utils.DECODED_PATIENT_ACTIVITY_DATA, &patientActivityData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
