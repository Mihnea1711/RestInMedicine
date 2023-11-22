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

	// Use dc.DB to save the doctor to the database
	lastInsertID, err := dController.DbConn.SaveDoctor(ctx, doctor)
	if err != nil {
		// Check if the error is due to a duplicate entry violation
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

		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[DOCTOR] Failed to save doctor to the database: %s\n", errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
			Error:   errMsg,
			Message: "Failed to create doctor.",
		})
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

	log.Printf("[DOCTOR] Successfully created doctor %d", lastInsertID)
	// Use RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: "Doctor created successfully",
		Payload: models.LastInsertedID{
			LastInsertedID: lastInsertID,
		},
	})
}
