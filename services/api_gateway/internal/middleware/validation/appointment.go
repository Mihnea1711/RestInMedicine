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

// ValidateAppointmentData validates the AppointmentData struct using the validator package
func validateAppointmentData(appointmentData models.AppointmentData) error {
	validate := validator.New()
	return validate.Struct(appointmentData)
}

// ValidateAppointmentData is a middleware that validates AppointmentData
func ValidateAppointmentData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var appointmentData models.AppointmentData

		contentTypeFlag := isContentTypeJSON(r)
		if !contentTypeFlag {
			errMsg := "Unsupported media type. Content-Type must be application/json"
			log.Printf("[MIDDLEWARE_GATEWAY] %s in request: %s", errMsg, r.RequestURI)
			utils.RespondWithJSON(w, http.StatusUnsupportedMediaType, models.ResponseData{Error: errMsg, Message: "Patient validation failed due to unsupported media type"})
			return
		}

		// Decode the request body into AppointmentData
		err := json.NewDecoder(r.Body).Decode(&appointmentData)
		if err != nil {
			logAndRespondWithError(w, http.StatusUnprocessableEntity, "Error decoding appointment request body", err)
			return
		}

		// Validate AppointmentData
		if err := validateAppointmentData(appointmentData); err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Validation error for appointment struct", err)
			return
		}

		// If validation passes, proceed to the next handler
		ctx := context.WithValue(r.Context(), utils.DECODED_APPOINTMENT_DATA, &appointmentData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
