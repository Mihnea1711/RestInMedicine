package controllers

import (
	"context"
	"database/sql"
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

	pController.handleContextTimeout(ctx, w)

	rowsAffected, err := pController.DbConn.DeletePatientByID(ctx, patientID)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Failed to delete patient with ID %d: %s", patientID, err)
			log.Printf("[PATIENT] %s", errMsg)

			// Use utils.RespondWithJSON for error response
			utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to delete patient. Patient not found"})
			return
		}

		errMsg := fmt.Sprintf("Failed to delete patient with ID %d: %s", patientID, err)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Failed to delete patient"})
		return
	}

	// Check if the patient exists and was deleted
	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("No patient found with ID: %d", patientID)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Patient not found or an unexpected error happened."})
		return
	}

	log.Printf("[PATIENT] Successfully deleted patient with ID %d", patientID)
	// Use utils.RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: fmt.Sprintf("Patient with ID: %d deleted successfully", patientID),
		Payload: models.RowsAffected{
			RowsAffected: rowsAffected,
		},
	})
}

func (pController *PatientController) DeletePatientByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientUserIDStr := vars[utils.DELETE_PATIENT_BY_USER_ID_PARAMETER]
	patientUserID, err := strconv.Atoi(patientUserIDStr)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid patient user ID: %s", patientUserIDStr)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Invalid patient delete request"})
		return
	}

	log.Printf("[PATIENT] Attempting to delete patient with user ID: %d...", patientUserID)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// getPatientBYUserID
	patient, err := pController.DbConn.FetchPatientByUserID(ctx, patientUserID)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Failed to fetch patient with user ID %d: %v", patientUserID, err)
			log.Printf("[PATIENT] %s", errMsg)

			// Use utils.RespondWithJSON for error response
			utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to fetch patient by user ID. Patient not found"})
			return
		}

		errMsg := fmt.Sprintf("Failed to fetch patient with user ID %d: %s", patientUserID, err)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Failed to fetch patient by patientUserID"})
		return
	}

	if patient == nil {
		errMsg := fmt.Sprintf("No patient found with user ID: %d", patientUserID)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Patient not found or an unexpected error happened."})
		return
	}

	rowsAffected, err := pController.DbConn.DeletePatientByUserID(ctx, patientUserID)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Failed to delete patient with user ID %d: %s", patientUserID, err)
			log.Printf("[PATIENT] %s", errMsg)

			// Use utils.RespondWithJSON for error response
			utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to delete patient. Patient not found"})
			return
		}

		errMsg := fmt.Sprintf("Failed to delete patient with user ID %d: %s", patientUserID, err)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Failed to delete patient"})
		return
	}

	// Check if the patient exists and was deleted
	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("No patient found with user ID: %d", patientUserID)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Patient not found or an unexpected error happened."})
		return
	}

	log.Printf("[PATIENT] Successfully deleted patient with user ID %d", patientUserID)
	// Use utils.RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: fmt.Sprintf("Patient with user ID: %d deleted successfully", patientUserID),
		Payload: models.ComplexResponse{
			RowsAffected: rowsAffected,
			DeletedID:    patient.IDPatient,
		},
	})
}
