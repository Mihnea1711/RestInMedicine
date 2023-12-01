package controllers

import (
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (dController *DoctorController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Printf("[DOCTOR] Handling health check request from %s\n", r.RemoteAddr)
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: "Health Check Handled Successfully",
	})
}
