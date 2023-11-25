package routes

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/idm"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/middleware"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/config"
)

func SetupRoutes(idmClient idm.IDMClient, jwtConfig config.JWTConfig) *mux.Router {
	log.Println("[GATEWAY] Setting up routes...")
	router := mux.NewRouter()
	router.Use(middleware.RouteLogger)

	gatewayController := &controllers.GatewayController{
		IDMClient: idmClient,
	}

	loadRoutes(router, gatewayController, jwtConfig)

	log.Println("[GATEWAY] Routes setup completed.")
	return router
}

func loadRoutes(router *mux.Router, gatewayController *controllers.GatewayController, jwtConfig config.JWTConfig) {
	log.Println("[GATEWAY] Loading routes for GATEWAY entity...")
	loadUserRoutes(router, gatewayController, jwtConfig)
	loadPatientRoutes(router, gatewayController, jwtConfig)
	loadDoctorRoutes(router, gatewayController, jwtConfig)
	loadAppointmentRoutes(router, gatewayController, jwtConfig)
	loadConsultationRoutes(router, gatewayController, jwtConfig)
	log.Println("[GATEWAY] All routes for GATEWAY entity loaded successfully.")
}
