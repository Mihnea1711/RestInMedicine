package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (dController *DoctorController) UpdateDoctorByID(w http.ResponseWriter, r *http.Request) {
	log.Println("[DOCTOR] Attempting to update a doctor.")

	// Decode the doctor details from the context
	doctor := r.Context().Value(utils.DECODED_DOCTOR).(*models.Doctor)

	// Get the doctor ID from the request path
	vars := mux.Vars(r)
	doctorIDStr := vars[utils.UPDATE_DOCTOR_BY_ID_PARAMETER]
	doctorID, err := strconv.Atoi(doctorIDStr)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid doctor ID: %s", doctorIDStr)
		log.Println("[DOCTOR] " + errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Invalid doctor update request"})
		return
	}

	// Assign the ID to the doctor object
	doctor.IDDoctor = doctorID

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	dController.handleContextTimeout(ctx, w)

	// Use dController.DbConn to update the doctor in the database
	rowsAffected, err := dController.DbConn.UpdateDoctorByID(ctx, doctor)
	if err != nil {
		handleDatabaseUpdateError(w, err)
		return
	}

	// Check if any rows were updated
	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("No doctor found with ID: %d", doctor.IDDoctor)
		log.Println("[DOCTOR] " + errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Doctor not found or an unexpected error happened."})
		return
	}

	log.Printf("[DOCTOR] Successfully updated doctor %d", doctor.IDDoctor)
	// Use RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: fmt.Sprintf("Doctor with ID: %d updated successfully", doctorID),
		Payload: models.RowsAffected{
			RowsAffected: rowsAffected,
		},
	})
}

func handleDatabaseUpdateError(w http.ResponseWriter, err error) {
	// Check if the error is due to no rows found
	if err == sql.ErrNoRows {
		errMsg := fmt.Sprintf("Error updating doctor: %s", err.Error())
		log.Printf("[DOCTOR] %s", errMsg)
		// Create a conflict response using ResponseData
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to update doctor. Doctor not found"})
		return
	}
	// Check if the error is a MySQL duplicate entry error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == utils.MySQLDuplicateEntryErrorCode {
		errMsg := fmt.Sprintf("Conflict error: %s", mysqlErr.Message)
		log.Printf("[DOCTOR] %s", errMsg)

		// Create a conflict response using ResponseData
		utils.RespondWithJSON(w, http.StatusConflict, models.ResponseData{Error: errMsg, Message: "Failed to update doctor. Duplicate entry violation"})
		return
	}

	errMsg := fmt.Sprintf("internal server error: %s", err)
	log.Printf("[DOCTOR] Failed to update doctor in the database: %s\n", errMsg)

	// Use RespondWithJSON for error response
	utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Internal database server error"})
}
