package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (dController *DoctorController) DeleteDoctorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid doctor ID: %s", idStr)
		log.Println("[DOCTOR] " + errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": errMsg})
		return
	}

	log.Printf("[DOCTOR] Attempting to delete doctor with ID: %d...", id)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	rowsAffected, err := dController.DbConn.DeleteDoctorByID(ctx, id)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to delete doctor with ID %d: %s", id, err)
		log.Println("[DOCTOR] " + errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": errMsg})
		return
	}

	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("No doctor found with ID: %d", id)
		log.Println("[DOCTOR] " + errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": errMsg})
		return
	}

	log.Printf("[DOCTOR] Successfully deleted doctor with ID: %d", id)

	// Use RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": fmt.Sprintf("Doctor with ID: %d deleted successfully", id)})
}
