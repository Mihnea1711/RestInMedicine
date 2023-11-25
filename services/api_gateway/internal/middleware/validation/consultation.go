package validation

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// ValidateConsultationData validates the ConsultationData struct using the validator package
func validateConsultationData(consultationData models.ConsultationData) error {
	validate := validator.New()
	return validate.Struct(consultationData)
}

// ValidateConsultationData is a middleware that validates ConsultationData
func ValidateConsultationData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var consultationData models.ConsultationData

		// Decode the request body into ConsultationData
		err := json.NewDecoder(r.Body).Decode(&consultationData)
		if err != nil {
			utils.SendErrorResponse(w, http.StatusBadRequest, "Error decoding consultation request body", err.Error())
			return
		}

		// Validate ConsultationData
		if err := validateConsultationData(consultationData); err != nil {
			utils.SendErrorResponse(w, http.StatusBadRequest, "Validation error for consultation struct", err.Error())
			return
		}

		// If validation passes, proceed to the next handler
		ctx := context.WithValue(r.Context(), utils.DECODED_CONSULTATION_DATA, &consultationData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
