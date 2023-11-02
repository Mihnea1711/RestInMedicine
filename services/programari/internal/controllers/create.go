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
	log.Printf("[PROGRAMARE] Attempting to create a new programare.")
	programare := r.Context().Value(utils.DECODED_PROGRAMARE).(*models.Programare)

	log.Printf("CREATE: %d / %d / %s / %s", programare.IDPacient, programare.IDDoctor, programare.Date.GoString(), programare.Status)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Use pController.DbConn to save the programare to the database
	err := pController.DbConn.SaveProgramare(ctx, programare)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[PROGRAMARE] Failed to save programare to the database: %s\n", errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	log.Printf("[PROGRAMARE] Successfully created programare %d", programare.IDProgramare)
	utils.RespondWithJSON(w, http.StatusOK, "Programare created")
}
