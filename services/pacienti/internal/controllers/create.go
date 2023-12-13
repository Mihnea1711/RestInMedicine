package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (pController *PatientController) CreatePatient(w http.ResponseWriter, r *http.Request) {
	log.Printf("[PATIENT] Attempting to create a new PATIENT.")
	patient := r.Context().Value(utils.DECODED_PATIENT).(*models.Patient)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	pController.handleContextTimeout(ctx, w)

	// Use dController.DbConn to save the patient to the database
	lastInsertID, err := pController.DbConn.SavePatient(ctx, patient)
	if err != nil {
		handleDatabaseCreateError(w, err)
		return
	}

	if lastInsertID == 0 {
		errorMsg := "Patient has not been saved to the database due to an unexpected error."
		log.Printf("[PATIENT] %s", errorMsg)

		// Use RespondWithJSON for conflict response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
			Message: "Failed to create patient",
			Error:   errorMsg,
		})
		return
	}

	// Construct the URI of the newly created patient
	baseURI := utils.FETCH_ALL_PATIENTS_ENDPOINT
	newPatientURI := fmt.Sprintf("%s/%d", baseURI, lastInsertID)

	log.Printf("[PATIENT] Successfully created patient %d", lastInsertID)

	// the origin server SHOULD send a 201 (Created) response containing a Location header field that provides an identifier for the primary resource created
	w.Header().Set("Location", newPatientURI)
	// Use RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusCreated, models.ResponseData{
		Message: "Patient created successfully",
		Payload: models.LastInsertedID{
			LastInsertedID: lastInsertID,
		},
	})
}

func handleDatabaseCreateError(w http.ResponseWriter, err error) {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if ok && mysqlErr.Number == utils.MySQLDuplicateEntryErrorCode {
		errMsg := fmt.Sprintf("Conflict error: %s", mysqlErr.Message)
		log.Printf("[PATIENT] %s", errMsg)

		// Create a conflict response using ResponseData
		utils.RespondWithJSON(w, http.StatusConflict, models.ResponseData{
			Error:   errMsg,
			Message: "Failed to create patient. Duplicate entry violation"})
		return
	}

	errMsg := fmt.Sprintf("Internal server error: %s", err)
	log.Printf("[PATIENT] Failed to save patient to the database: %s\n", errMsg)

	// Use RespondWithJSON for error response
	utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
		Message: "Failed to create patient",
		Error:   errMsg,
	})
}
