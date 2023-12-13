package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (dController *DoctorController) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	log.Printf("[DOCTOR] Attempting to create a new doctor.")
	doctor := r.Context().Value(utils.DECODED_DOCTOR).(*models.Doctor)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	dController.handleContextTimeout(ctx, w)

	// Use dc.DB to save the doctor to the database
	lastInsertID, err := dController.DbConn.SaveDoctor(ctx, doctor)
	if err != nil {
		handleDatabaseCreateError(w, err)
		return
	}

	if lastInsertID == 0 {
		errorMsg := "Dcoctor has not been saved to the database due to an unexpected error."
		log.Printf("[DOCTOR] %s", errorMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
			Message: "Failed to create doctor",
			Error:   errorMsg,
		})
		return
	}

	// Construct the URI of the newly created patient
	baseURI := utils.FETCH_ALL_DOCTORS_ENDPOINT
	newDoctorURI := fmt.Sprintf("%s/%d", baseURI, lastInsertID)

	log.Printf("[DOCTOR] Successfully created doctor %d", lastInsertID)

	// the origin server SHOULD send a 201 (Created) response containing a Location header field that provides an identifier for the primary resource created
	w.Header().Set("Location", newDoctorURI)
	// Use RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusCreated, models.ResponseData{
		Message: "Doctor created successfully",
		Payload: models.LastInsertedID{
			LastInsertedID: lastInsertID,
		},
	})
}

func handleDatabaseCreateError(w http.ResponseWriter, err error) {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if ok && mysqlErr.Number == utils.MySQLDuplicateEntryErrorCode {
		errMsg := fmt.Sprintf("Conflict error: %s", mysqlErr.Message)
		log.Printf("[DOCTOR] %s", errMsg)

		// Create a conflict response using ResponseData
		utils.RespondWithJSON(w, http.StatusConflict, models.ResponseData{
			Error:   errMsg,
			Message: "Failed to create doctor. Duplicate entry violation"})
		return
	}

	errMsg := fmt.Sprintf("Internal server error: %s", err)
	log.Printf("[DOCTOR] Failed to save doctor to the database: %s\n", errMsg)

	// Use RespondWithJSON for error response
	utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
		Message: "Failed to create doctor",
		Error:   errMsg,
	})
}
