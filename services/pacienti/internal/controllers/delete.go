package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (pController *PatientController) DeletePatientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientIDStr := vars[utils.DELETE_PATIENT_BY_ID_PARAMETER]

	patientID, err := strconv.Atoi(patientIDStr)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid patient ID: %s", patientIDStr)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Invalid patient delete request"})
		return
	}

	log.Printf("[PATIENT] Attempting to delete patient with ID: %d...", patientID)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	rowsAffected, err := pController.DbConn.DeletePatientByID(ctx, patientID)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to delete patient with ID %d: %s", patientID, err)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Failed to delete patient"})
		return
	}

	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("No patient found with ID: %d", patientID)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Patient not found or an unexpected error happened."})
		return
	}

	log.Printf("[PATIENT] Successfully deleted patient with ID %d", patientID)
	// Use utils.RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{Message: fmt.Sprintf("Patient with ID: %d deleted successfully", patientID)})
}
