package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// CreateAppointment handles the creation of a new programare.
func (c *GatewayController) CreateAppointment(w http.ResponseWriter, r *http.Request) {
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
	response, err := c.redirectRequestBody(utils.POST, utils.CREATE_APPOINTMENT_ENDPOINT, utils.APPOINTMENT_PORT, appointmentRequest)
	if err != nil {
		// Handle the error (e.g., return a response with an error message)
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to redirect request"})
		return
	}

	// Respond with the response from the other module
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"message": "Appointment made successfully", "app_data": appointmentRequest, "response": response})
}

// GetAppointments handles the retrieval of all programari.
func (c *GatewayController) GetAppointments(w http.ResponseWriter, r *http.Request) {

}

// GetAppointmentByID handles the retrieval of a programare by ID.
func (c *GatewayController) GetAppointmentByID(w http.ResponseWriter, r *http.Request) {

}

// GetAppointmentsByDoctorID handles the retrieval of programari by doctor ID.
func (c *GatewayController) GetAppointmentsByDoctorID(w http.ResponseWriter, r *http.Request) {

}

// GetAppointmentsByPacientID handles the retrieval of programari by pacient ID.
func (c *GatewayController) GetAppointmentsByPacientID(w http.ResponseWriter, r *http.Request) {

}

// GetAppointmentsByDate handles the retrieval of programari by date.
func (c *GatewayController) GetAppointmentsByDate(w http.ResponseWriter, r *http.Request) {

}

// GetAppointmentsByStatus handles the retrieval of programari by status.
func (c *GatewayController) GetAppointmentsByStatus(w http.ResponseWriter, r *http.Request) {

}

// UpdateAppointmentByID handles the update of a specific programare by ID.
func (c *GatewayController) UpdateAppointmentByID(w http.ResponseWriter, r *http.Request) {

}

// DeleteAppointmentByID handles the deletion of a programare by ID.
func (c *GatewayController) DeleteAppointmentByID(w http.ResponseWriter, r *http.Request) {

}
