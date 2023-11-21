package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

// Delete a appointment by ID
func (pController *ProgramareController) DeleteProgramareByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[APPOINTMENT] Attempting to delete a appointment by ID.")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars[utils.DELETE_APPOINTMENT_BY_ID_PARAMETER])
	if err != nil {
		response := models.ResponseData{Error: "Invalid appointment ID"}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Use pController.DbConn to delete the appointment by ID from the database
	rowsAffected, err := pController.DbConn.DeleteProgramareByID(ctx, id)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[APPOINTMENT] Failed to delete appointment by ID: %s\n", errMsg)
		response := models.ResponseData{Error: errMsg}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Check if the appointment exists and was deleted
	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("No appointment found with ID: %d", id)
		log.Println("[APPOINTMENT] " + errMsg)

		response := models.ResponseData{Error: "Appointment not found"}
		utils.RespondWithJSON(w, http.StatusNotFound, response)
		return
	}

	log.Printf("[APPOINTMENT] Successfully deleted appointment %d", id)
	response := models.ResponseData{Message: "Appointment deleted"}
	utils.RespondWithJSON(w, http.StatusOK, response)
}
