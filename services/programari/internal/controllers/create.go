package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

func (pController *ProgramareController) CreateProgramare(w http.ResponseWriter, r *http.Request) {
	log.Printf("[APPOINTMENT] Attempting to create a new appointment.")
	programare := r.Context().Value(utils.DECODED_APPOINTMENT).(*models.Programare)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Use pController.DbConn to save the appointment to the database
	lastInsertID, err := pController.DbConn.SaveProgramare(ctx, programare)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[APPOINTMENT] Failed to save programare to the database: %s\n", errMsg)
		response := models.ResponseData{Error: errMsg}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	if lastInsertID == 0 {
		utils.RespondWithJSON(w, http.StatusConflict, models.ResponseData{
			Message: "Appointment not saved",
		})
		return
	}

	log.Printf("[APPOINTMENT] Successfully created programare %d", lastInsertID)
	response := models.ResponseData{Message: "Programare created"}
	utils.RespondWithJSON(w, http.StatusOK, response)
}
