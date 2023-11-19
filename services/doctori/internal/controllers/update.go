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
	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (dController *DoctorController) UpdateDoctorByID(w http.ResponseWriter, r *http.Request) {
	log.Println("[DOCTOR] Attempting to update a doctor.")

	// Decode the doctor details from the context (assuming you've set it in the middleware)
	doctor := r.Context().Value(utils.DECODED_DOCTOR).(*models.Doctor)

	// Get the doctor ID from the request path
	vars := mux.Vars(r)
	idStr := vars[utils.UPDATE_DOCTOR_BY_ID_PARAMETER]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid doctor ID: %s", idStr)
		log.Println("[DOCTOR] " + errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg})
		return
	}

	// Assign the ID to the doctor object
	doctor.IDDoctor = id

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Use dController.DbConn to update the doctor in the database
	rowsAffected, err := dController.DbConn.UpdateDoctorByID(ctx, doctor)
	if err != nil {
		// Check if the error is a MySQL duplicate entry error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == utils.MySQLDuplicateEntryErrorCode {
			errMsg := fmt.Sprintf("Conflict error: %s", mysqlErr.Message)
			log.Printf("[DOCTOR] %s", errMsg)

			// Create a conflict response using ResponseData
			utils.RespondWithJSON(w, http.StatusConflict, models.ResponseData{Error: errMsg})
			return
		}

		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[DOCTOR] Failed to update doctor in the database: %s\n", errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg})
		return
	}

	// Check if any rows were updated
	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("No doctor found with ID: %d", doctor.IDDoctor)
		log.Println("[DOCTOR] " + errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg})
		return
	}

	log.Printf("[DOCTOR] Successfully updated doctor %d", doctor.IDDoctor)

	// Use RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{Message: "Doctor updated"})
}
