package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (dController *PacientController) DeletePacientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		err_msg := fmt.Sprintf("Invalid pacient ID: %s", idStr)
		log.Printf("[PACIENT] %s", err_msg)
		http.Error(w, err_msg, http.StatusBadRequest)
		return
	}

	log.Printf("[PACIENT] Attempting to delete pacient with ID: %d...", id)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	rowsAffected, err := dController.DbConn.DeletePacientByID(ctx, id)
	if err != nil {
		err_msg := fmt.Sprintf("Failed to delete pacient with ID %d: %s", id, err)
		log.Printf("[PACIENT] %s", err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		err_msg := fmt.Sprintf("No pacient found with ID: %d", id)
		log.Printf("[PACIENT] %s", err_msg)
		http.Error(w, err_msg, http.StatusNotFound)
		return
	}

	log.Printf("[PACIENT] Successfully deleted pacient with ID: %d", id)
	w.Write([]byte(fmt.Sprintf("Pacient with ID: %d deleted successfully\n", id)))
}
