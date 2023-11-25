package validation

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// ValidateDoctorData validates the DoctorData struct using the validator package
func validateDoctorData(doctorData models.DoctorData) error {
	validate := validator.New()
	return validate.Struct(doctorData)
}

// ValidateDoctorData is a middleware that validates DoctorData
func ValidateDoctorData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var doctorData models.DoctorData

		// Decode the request body into DoctorData
		err := json.NewDecoder(r.Body).Decode(&doctorData)
		if err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Error decoding doctor request body", err)
			return
		}

		// Validate DoctorData
		if err := validateDoctorData(doctorData); err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Validation error for doctor struct", err)
			return
		}

		// If validation passes, proceed to the next handler
		ctx := context.WithValue(r.Context(), utils.DECODED_DOCTOR_DATA, &doctorData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
