package controllers

import (
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

func (aController *AppointmentController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Printf("[APPOINTMENT] Handling health check request from %s\n", r.RemoteAddr)
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: "Health Check Handled Successfully",
	})
}
