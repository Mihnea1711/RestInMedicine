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
	//

	// Redirect the request body to another module
	response, err := c.redirectRequestBody(utils.CREATE_APPOINTMENT_ENDPOINT, utils.APPOINTMENT_PORT, appointmentRequest)
	if err != nil {
		// Handle the error (e.g., return a response with an error message)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to redirect request"})
		return
	}

	// Respond with the response from the other module
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"message": "Appointment made successfully", "app_data": appointmentRequest, "response": response})
}
