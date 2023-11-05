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
	log.Println("[PACIENT] Attempting to update a pacient.")

	// Decode the pacient details from the context (assuming you've set it in the middleware)
	pacient := r.Context().Value(utils.DECODED_PACIENT).(*models.Pacient)

	// Get the pacient ID from the request path
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err_msg := fmt.Sprintf("Invalid pacient ID: %s", idStr)
		log.Printf("[PACIENT] %s", err_msg)
		http.Error(w, err_msg, http.StatusBadRequest)
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
		err_msg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[PACIENT] Failed to update pacient in the database: %s\n", err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	// Check if any rows were updated
	if rowsAffected == 0 {
		err_msg := fmt.Sprintf("No pacient found with ID: %d", pacient.IDPacient)
		log.Printf("[PACIENT] %s", err_msg)

		http.Error(w, err_msg, http.StatusNotFound)
		return
	}

	log.Printf("[PACIENT] Successfully updated pacient %d", pacient.IDPacient)
	w.Write([]byte("Pacient updated\n"))
}
