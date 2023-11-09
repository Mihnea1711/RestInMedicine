package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/controllers"
)

func loadAppointmentRoutes(router *mux.Router, gatewayController *controllers.GatewayController) {
	// MakeAppointment handles the creation of a new appointment.
	makeAppointmentHandler := http.HandlerFunc(gatewayController.MakeAppointment)
	router.Handle("/api/appointments", makeAppointmentHandler).Methods("POST")
	log.Println("[GATEWAY] Route POST /api/make-appointment registered.")
}
