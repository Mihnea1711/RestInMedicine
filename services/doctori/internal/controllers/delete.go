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

func (dController *DoctorController) DeleteDoctorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		err_msg := fmt.Sprintf("Invalid doctor ID: %s", idStr)
		log.Println("[DOCTOR] " + err_msg)
		http.Error(w, err_msg, http.StatusBadRequest)
		return
	}

	log.Printf("[DOCTOR] Attempting to delete doctor with ID: %d...", id)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	rowsAffected, err := dController.DbConn.DeleteDoctorByID(ctx, id)
	if err != nil {
		err_msg := fmt.Sprintf("Failed to delete doctor with ID %d: %s", id, err)
		log.Println("[DOCTOR] " + err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		err_msg := fmt.Sprintf("No doctor found with ID: %d", id)
		log.Println("[DOCTOR] " + err_msg)
		http.Error(w, err_msg, http.StatusNotFound)
		return
	}

	log.Printf("[DOCTOR] Successfully deleted doctor with ID: %d", id)
	w.Write([]byte(fmt.Sprintf("Doctor with ID: %d deleted successfully\n", id)))
}
