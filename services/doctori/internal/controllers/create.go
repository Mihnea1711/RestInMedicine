package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (dController *DoctorController) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	log.Printf("[DOCTOR] Attempting to create a new doctor.")
	doctor := r.Context().Value(utils.DECODED_DOCTOR).(*models.Doctor)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Use dc.DB to save the doctor to the database
	err := dController.DbConn.SaveDoctor(ctx, doctor)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[DOCTOR] Failed to save doctor to the database: %s\n", errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
			Error: errMsg,
		})
		return
	}

	log.Printf("[DOCTOR] Successfully created doctor %d", doctor.IDDoctor)

	// Use RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: "Doctor created",
	})
}
