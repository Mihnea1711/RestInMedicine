package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/idm"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/middleware"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/middleware/validation"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/config"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

func SetupRoutes(idmClient idm.IDMClient, jwtConfig config.JWTConfig) *mux.Router {
	router := mux.NewRouter()
	log.Println("[GATEWAY] Setting up routes...")

	router.Use(middleware.RouteLogger)
	log.Println("[GATEWAY] Route logger middleware set up successfully.")

	router.Use(middleware.SanitizeInputMiddleware)
	log.Println("[GATEWAY] Input sanitizer middleware set up successfully.")

	// register custom validations
	validation.RegisterCustomValidationTags()

	gatewayController := &controllers.GatewayController{
		IDMClient: idmClient,
	}

	loadRoutes(router, gatewayController, jwtConfig)

	router.Use(middleware.AddPathAndMethodToResponse)

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
	loadHealthRoutes(router, gatewayController)
	log.Println("[GATEWAY] All routes for GATEWAY entity loaded successfully.")
}

// loadUserRoutes loads all the CRUD routes for the User entity
func loadHealthRoutes(router *mux.Router, gatewayController *controllers.GatewayController) {
	healthHandler := http.HandlerFunc(gatewayController.HealthCheck)
	router.Handle(utils.CHECK_HEALTH_ENDPOINT, healthHandler).Methods("GET")
	log.Println("[GATEWAY] Route POST", utils.CHECK_HEALTH_ENDPOINT, "registered.")
}
