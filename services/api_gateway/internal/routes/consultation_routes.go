package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/controllers"
)

func loadConsultationRoutes(router *mux.Router, gatewayController *controllers.GatewayController) {
	// ProgramConsultation handles the scheduling of a new consultation.
	programConsultationHandler := http.HandlerFunc(gatewayController.ProgramConsultation)
	router.Handle("/api/consultations", programConsultationHandler).Methods("POST")
	log.Println("[GATEWAY] Route POST /api/program-consultation registered.")
}
