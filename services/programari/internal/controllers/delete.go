package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

// Delete a programare by ID
func (pController *ProgramareController) DeleteProgramareByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[APPOINTMENT] Attempting to delete a programare by ID.")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid programare ID", http.StatusBadRequest)
		return
	}

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Use pController.DbConn to delete the programare by ID from the database
	rowsAffected, err := pController.DbConn.DeleteProgramareByID(ctx, id)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[APPOINTMENT] Failed to delete programare by ID: %s\n", errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	// Check if the programare exists and was deleted
	if rowsAffected == 0 {
		http.Error(w, "Programare not found", http.StatusNotFound)
		return
	}

	log.Printf("[APPOINTMENT] Successfully deleted programare %d", id)
	utils.RespondWithJSON(w, http.StatusOK, "Programare deleted")
}
