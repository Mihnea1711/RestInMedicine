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
	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

// Update a programare by ID
func (pController *ProgramareController) UpdateProgramareByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[APPOINTMENT] Attempting to update an appointment by ID.")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars[utils.UPDATE_APPOINTMENT_BY_ID_PARAMETER])
	if err != nil {
		response := models.ResponseData{Error: "Invalid appointment ID"}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	programare := r.Context().Value(utils.DECODED_APPOINTMENT).(*models.Programare)
	programare.IDProgramare = id

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Use pController.DbConn to update the appointment by ID in the database
	rowsAffected, err := pController.DbConn.UpdateProgramareByID(ctx, programare)
	if err != nil {
		// Check if the error is a MySQL duplicate entry error
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == utils.MySQLDuplicateEntryErrorCode {
			errMsg := fmt.Sprintf("Conflict error: %s", mysqlErr.Message)
			log.Printf("[APPOINTMENT] %s", errMsg)

			// Create a conflict response using ResponseData
			utils.RespondWithJSON(w, http.StatusConflict, models.ResponseData{Error: errMsg})
			return
		}

		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[APPOINTMENT] Failed to update appointment by ID: %s\n", errMsg)
		response := models.ResponseData{Error: errMsg}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Check if the appointment exists and was updated
	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("No programare found with ID: %d", programare.IDProgramare)
		log.Println("[APPOINTMENT] " + errMsg)

		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: "Programare not found"})
		return
	}

	log.Printf("[APPOINTMENT] Successfully updated appointment %d", programare.IDProgramare)
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{Message: "Appointment updated"})
}
