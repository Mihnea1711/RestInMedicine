package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (dController *DoctorController) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	doctor := r.Context().Value(utils.DECODED_DOCTOR).(*models.Doctor)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Use dc.DB to save the doctor to the database
	err := dController.DbConn.SaveDoctor(ctx, doctor)
	if err != nil {
		err_msg := fmt.Sprintf("internal server error: %s", err)
		http.Error(w, err_msg, http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Doctor created\n"))
}
