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

func (dController *PacientController) UpdatePacientByID(w http.ResponseWriter, r *http.Request) {
	log.Println("[PATIENT] Attempting to update a pacient.")

	// Decode the pacient details from the context (assuming you've set it in the middleware)
	pacient := r.Context().Value(utils.DECODED_PATIENT).(*models.Pacient)

	// Get the pacient ID from the request path
	vars := mux.Vars(r)
	idStr := vars[utils.UPDATE_PATIENT_BY_ID_PARAMETER]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid pacient ID: %s", idStr)
		log.Printf("[PATIENT] %s", errMsg)

		// Create an error response using ResponseData
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg})
		return
	}

	// Assign the ID to the pacient object
	pacient.IDPacient = id

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Use dController.DbConn to update the pacient in the database
	rowsAffected, err := dController.DbConn.UpdatePacientByID(ctx, pacient)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[PATIENT] Failed to update pacient in the database: %s\n", errMsg)

		// Create an error response using ResponseData
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg})
		return
	}

	// Check if any rows were updated
	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("No pacient found with ID: %d", pacient.IDPacient)
		log.Printf("[PATIENT] %s", errMsg)

		// Create an error response using ResponseData
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg})
		return
	}

	// Create a success response using ResponseData
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{Message: "Pacient updated"})
}
