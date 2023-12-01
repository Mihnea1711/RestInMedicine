package controllers

import (
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
)

func (cController *ConsultationController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Printf("[CONSULTATION] Handling health check request from %s\n", r.RemoteAddr)
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: "Health Check Handled Successfully",
	})
}
