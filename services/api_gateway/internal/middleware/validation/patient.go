package validation

import (
	"context"
	"encoding/json"
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

		// Decode the request body into PatientData
		err := json.NewDecoder(r.Body).Decode(&patientData)
		if err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Error decoding patient request body", err)
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
