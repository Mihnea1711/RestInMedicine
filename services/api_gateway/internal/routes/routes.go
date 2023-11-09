package routes

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/middleware"
)

func SetupRoutes() *mux.Router {
	log.Println("[GATEWAY] Setting up routes...")
	router := mux.NewRouter()
	router.Use(middleware.RouteLogger)

	gatewayController := &controllers.GatewayController{}

	loadRoutes(router, gatewayController)

	log.Println("[GATEWAY] Routes setup completed.")
	return router
}

func loadRoutes(router *mux.Router, gatewayController *controllers.GatewayController) {
	log.Println("[GATEWAY] Loading routes for GATEWAY entity...")
	loadUserRoutes(router, gatewayController)
	loadPatientRoutes(router, gatewayController)
	loadDoctorRoutes(router, gatewayController)
	loadAppointmentRoutes(router, gatewayController)
	loadConsultationRoutes(router, gatewayController)
	log.Println("[GATEWAY] All routes for GATEWAY entity loaded successfully.")
}
