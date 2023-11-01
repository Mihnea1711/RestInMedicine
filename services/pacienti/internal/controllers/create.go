package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (dController *PacientController) CreatePacient(w http.ResponseWriter, r *http.Request) {
	log.Printf("[PACIENTI] Attempting to create a new PACIENT.")
	doctor := r.Context().Value(utils.DECODED_PACIENT).(*models.Pacient)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Use dc.DB to save the doctor to the database
	err := dController.DbConn.SavePacient(ctx, doctor)
	if err != nil {
		err_msg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[PACIENTI] Failed to save doctor to the database: %s\n", err_msg)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}

	log.Printf("[PACIENTI] Successfully created doctor %d", doctor.IDPacient)
	w.Write([]byte("Doctor created\n"))
}
