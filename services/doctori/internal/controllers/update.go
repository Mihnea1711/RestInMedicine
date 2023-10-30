package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (dController *DoctorController) UpdateDoctorByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to update a doctor.")

	// Decode the doctor details from the context (assuming you've set it in the middleware)
	doctor := r.Context().Value(utils.DECODED_DOCTOR).(*models.Doctor)

	// Get the doctor ID from the request path
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err_msg := fmt.Sprintf("Invalid doctor ID: %s", idStr)
		log.Println(err_msg)
		http.Error(w, err_msg, http.StatusBadRequest)
		return
	}

	// Assign the ID to the doctor object
	doctor.IDDoctor = id

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Use dController.DbConn to update the doctor in the database
	rowsAffected, err := dController.DbConn.UpdateDoctorByID(ctx, doctor)
	if err != nil {
		err_msg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("Failed to update doctor in the database: %s\n", err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	// Check if any rows were updated
	if rowsAffected == 0 {
		err_msg := fmt.Sprintf("No doctor found with ID: %d", doctor.IDDoctor)
		log.Println(err_msg)
		http.Error(w, err_msg, http.StatusNotFound)
		return
	}

	log.Printf("Successfully updated doctor %d", doctor.IDDoctor)
	w.Write([]byte("Doctor updated\n"))
}
