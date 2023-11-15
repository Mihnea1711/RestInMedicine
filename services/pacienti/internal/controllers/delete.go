package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (dController *PacientController) DeletePacientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid pacient ID: %s", idStr)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": errMsg})
		return
	}

	log.Printf("[PATIENT] Attempting to delete pacient with ID: %d...", id)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	rowsAffected, err := dController.DbConn.DeletePacientByID(ctx, id)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to delete pacient with ID %d: %s", id, err)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": errMsg})
		return
	}

	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("No pacient found with ID: %d", id)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": errMsg})
		return
	}

	// Use utils.RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": fmt.Sprintf("Pacient with ID: %d deleted successfully", id)})
}
