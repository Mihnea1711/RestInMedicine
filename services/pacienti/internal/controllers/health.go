package controllers

import (
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (pController *PatientController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Printf("[PATIENT] Handling health check request from %s\n", r.RemoteAddr)
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: "Health Check Handled Successfully",
	})
}
