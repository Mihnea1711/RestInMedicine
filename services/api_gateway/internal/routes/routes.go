package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/middleware"
)

func SetupRoutes() *mux.Router {
	log.Println("[GATEWAY] Setting up routes...")
	router := mux.NewRouter()
	router.Use(middleware.RouteLogger)

	gatewayController := &controllers.GatewayController{}

	loadCrudRoutes(router, gatewayController)

	log.Println("[GATEWAY] Routes setup completed.")
	return router
}

// loadCrudRoutes loads all the CRUD routes for the GATEWAY entity
func loadCrudRoutes(router *mux.Router, gatewayController *controllers.GatewayController) {
	log.Println("[GATEWAY] Loading routes for GATEWAY entity...")

	// RegisterUser handles user registration.
	registerUserHandler := http.HandlerFunc(gatewayController.RegisterUser)
	router.Handle("/api/register", registerUserHandler).Methods("POST").Queries("role", "{role}")
	log.Println("[GATEWAY] Route POST /api/register registered.")

	// LoginUser handles user login.
	loginUserHandler := http.HandlerFunc(gatewayController.LoginUser)
	router.Handle("/api/login", loginUserHandler).Methods("POST")
	log.Println("[GATEWAY] Route POST /api/login registered.")

	// MakeAppointment handles the creation of a new appointment.
	makeAppointmentHandler := http.HandlerFunc(gatewayController.MakeAppointment)
	router.Handle("/api/appointments", makeAppointmentHandler).Methods("POST")
	log.Println("[GATEWAY] Route POST /api/make-appointment registered.")

	// ProgramConsultation handles the scheduling of a new consultation.
	programConsultationHandler := http.HandlerFunc(gatewayController.ProgramConsultation)
	router.Handle("/api/consultations", programConsultationHandler).Methods("POST")
	log.Println("[GATEWAY] Route POST /api/program-consultation registered.")

	log.Println("[GATEWAY] All routes for GATEWAY entity loaded successfully.")
}
