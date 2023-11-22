package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (pController *PatientController) UpdatePatientByID(w http.ResponseWriter, r *http.Request) {
	log.Println("[PATIENT] Attempting to update a patient.")

	// Decode the patient details from the context (assuming you've set it in the middleware)
	patient := r.Context().Value(utils.DECODED_PATIENT).(*models.Pacient)

	// Get the patient ID from the request path
	vars := mux.Vars(r)
	patientIDStr := vars[utils.UPDATE_PATIENT_BY_ID_PARAMETER]
	patientID, err := strconv.Atoi(patientIDStr)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid patient ID: %s", patientIDStr)
		log.Printf("[PATIENT] %s", errMsg)

		// Create an error response using ResponseData
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Invalid patient update request"})
		return
	}

	// Assign the ID to the patient object
	patient.IDPacient = patientID

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Use pController.DbConn to update the patient in the database
	rowsAffected, err := pController.DbConn.UpdatePatientByID(ctx, patient)
	if err != nil {
		// Check if the error is a MySQL duplicate entry error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == utils.MySQLDuplicateEntryErrorCode {
			errMsg := fmt.Sprintf("Conflict error: %s", mysqlErr.Message)
			log.Printf("[PATIENT] %s", errMsg)

			// Create a conflict response using ResponseData
			utils.RespondWithJSON(w, http.StatusConflict, models.ResponseData{Error: errMsg, Message: "Failed to update patient"})
			return
		}

		errMsg := fmt.Sprintf("Internal server error: %s", err)
		log.Printf("[PATIENT] Failed to update patient in the database: %s\n", errMsg)

		// Create an error response using ResponseData
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Internal database server error"})
		return
	}

	// Check if any rows were updated
	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("No patient found with ID: %d", patient.IDPacient)
		log.Printf("[PATIENT] %s", errMsg)

		// Create an error response using ResponseData
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Patient not found or an unexpected error happened."})
		return
	}

	log.Printf("[PATIENT] Successfully updated patient with ID %d", patient.IDPacient)
	// Create a success response using ResponseData
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{Message: "Patient updated successfully", Payload: patient})
}
