package controllers

import (
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

func (gc *GatewayController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Handling health check request from %s\n", r.RemoteAddr)
	utils.SendMessageResponse(w, http.StatusOK, "Health Check Handled Successfully", models.ResponseData{
		Message: "Health Check Handled Successfully",
	})
}

func (gc *GatewayController) TestHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Handling test check request from %s\n", r.RemoteAddr)

	log.Printf("[GATEWAY] Request Headers: %+v\n", r.Header)

	// Handle your request logic here
	utils.SendMessageResponse(w, http.StatusOK, "Test Check Handled Successfully", models.ResponseData{
		Message: "Test Check Handled Successfully",
	})
}
