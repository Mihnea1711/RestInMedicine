package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// ProgramConsultation handles programming a consultation.
func (c *GatewayController) ProgramConsultation(w http.ResponseWriter, r *http.Request) {
	var consultationRequest models.ProgramConsultationRequest

	// Parse the request body into the ProgramConsultationRequest struct
	if err := json.NewDecoder(r.Body).Decode(&consultationRequest); err != nil {
		// Handle the error (e.g., return a response with an error message)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	// Perform your consultation logic...

	// Respond with a success message
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"message": "Consultation programmed successfully", "cons_data": consultationRequest})
}
