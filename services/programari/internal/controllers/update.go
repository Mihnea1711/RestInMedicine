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

func (pController *ProgramareController) UpdateProgramareByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[PROGRAMARE] Attempting to update a programare by ID.")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid programare ID", http.StatusBadRequest)
		return
	}

	programare := r.Context().Value(utils.DECODED_PROGRAMARE).(*models.Programare)
	programare.IDProgramare = id

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Use pController.DbConn to update the programare by ID in the database
	rowsAffected, err := pController.DbConn.UpdateProgramareByID(ctx, programare)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[PROGRAMARE] Failed to update programare by ID: %s\n", errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	// Check if the programare exists and was updated
	if rowsAffected == 0 {
		http.Error(w, "Programare not found", http.StatusNotFound)
		return
	}

	log.Printf("[PROGRAMARE] Successfully updated programare %d", programare.IDProgramare)
	w.Write([]byte("Programare updated\n"))

	// utils.RespondWithJSON()
}
