package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/idm"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/middleware"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/middleware/authorization"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/middleware/cors_config"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/middleware/validation"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/config"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

func SetupRoutes(idmClient idm.IDMClient, jwtConfig config.JWTConfig) *mux.Router {
	router := mux.NewRouter()
	log.Println("[GATEWAY] Setting up routes...")

	cors_config.SetupCORS(router)
	log.Println("[GATEWAY] CORS middleware set up successfully.")

	router.Use(middleware.RouteLogger)
	log.Println("[GATEWAY] Route logger middleware set up successfully.")

	router.Use(middleware.SanitizeInputMiddleware)
	log.Println("[GATEWAY] Input sanitizer middleware set up successfully.")

	router.Use(authorization.BlacklistMiddleware(idmClient, jwtConfig))
	log.Println("[GATEWAY] Blacklist validator middleware set up successfully.")

	// register custom validations
	validation.RegisterCustomValidationTags()

	gatewayController := &controllers.GatewayController{
		IDMClient: idmClient,
	}

	loadRoutes(router, gatewayController, jwtConfig)

	router.Use(middleware.AddPathAndMethodToResponse)
	log.Println("[GATEWAY] HATEOAS bonus middleware set up successfully.")

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
	loadOtherRoutes(router, gatewayController)
	log.Println("[GATEWAY] All routes for GATEWAY entity loaded successfully.")
}

// loadUserRoutes loads all the CRUD routes for the User entity
func loadOtherRoutes(router *mux.Router, gatewayController *controllers.GatewayController) {
	healthHandler := http.HandlerFunc(gatewayController.HealthCheck)
	router.Handle(utils.CHECK_HEALTH_ENDPOINT, healthHandler).Methods("GET")
	log.Println("[GATEWAY] Route GET", utils.CHECK_HEALTH_ENDPOINT, "registered.")

	docsHandler := http.HandlerFunc(gatewayController.GetDocs)
	router.Handle(utils.GET_DOC_ENDPOINT, docsHandler).Methods("GET")
	log.Println("[GATEWAY] Route GET", utils.GET_DOC_ENDPOINT, "registered.")

	testHandler := http.HandlerFunc(gatewayController.TestHandler)
	router.Handle("/api/test", testHandler).Methods("POST")
	log.Println("[GATEWAY] Route POST", "/api/test", "registered.")
}
