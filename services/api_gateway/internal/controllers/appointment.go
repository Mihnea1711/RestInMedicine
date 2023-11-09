package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// MakeAppointment handles making an appointment.
func (c *GatewayController) MakeAppointment(w http.ResponseWriter, r *http.Request) {
	var appointmentRequest models.MakeAppointmentRequest

	// Parse the request body into the MakeAppointmentRequest struct
	if err := json.NewDecoder(r.Body).Decode(&appointmentRequest); err != nil {
		// Handle the error (e.g., return a response with an error message)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	// Perform your appointment logic...

	// Respond with a success message
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"message": "Appointment made successfully", "app_data": appointmentRequest})
}
