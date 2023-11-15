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
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Use pController.DbConn to save the appointment to the database
	err := pController.DbConn.SaveProgramare(ctx, programare)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[APPOINTMENT] Failed to save programare to the database: %s\n", errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	log.Printf("[APPOINTMENT] Successfully created programare %d", programare.IDProgramare)
	utils.RespondWithJSON(w, http.StatusOK, "Programare created")
}
