package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/middleware"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/middleware/authorization"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/middleware/cors_config"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/rabbitmq"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/config"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/utils"
)

func SetupRoutes(rabbitMQ *rabbitmq.RabbitMQ, jwtConfig config.JWTConfig) *mux.Router {
	log.Println("[RABBIT] Setting up routes...")
	router := mux.NewRouter()

	cors_config.SetupCORS(router)
	log.Println("[RABBIT] CORS middleware set up successfully.")

	router.Use(middleware.RouteLogger)
	log.Println("[RABBIT] Route logger middleware set up successfully.")

	router.Use(middleware.SanitizeInputMiddleware)
	log.Println("[RABBIT] Input sanitizer middleware set up successfully.")

	rabbitController := &controllers.RabbitController{
		RabbitMQ: rabbitMQ,
	}

	loadRoutes(router, rabbitController, jwtConfig)

	log.Println("[RABBIT] Routes setup completed.")
	return router
}

func loadRoutes(router *mux.Router, rabbitController *controllers.RabbitController, jwtConfig config.JWTConfig) {
	log.Println("[RABBIT] Loading routes for RABBIT entity...")

	// ---------------------------------------------------------- Health --------------------------------------------------------------
	healthCheckHandler := http.HandlerFunc(rabbitController.HandleHealthCheck)
	router.Handle(utils.HEALTH_CHECK_ENDPOINT, healthCheckHandler).Methods(http.MethodGet)
	log.Println("[RABBIT] Route GET", utils.HEALTH_CHECK_ENDPOINT, "registered.")

	// ---------------------------------------------------------- PUBLISH --------------------------------------------------------------
	publishHandler := http.HandlerFunc(rabbitController.Publish)
	router.Handle(utils.PUBLISH_ENDPOINT, authorization.AdminOnlyMiddleware(jwtConfig, publishHandler)).Methods(http.MethodPost)
	log.Println("[RABBIT] Route POST", utils.PUBLISH_ENDPOINT, "registered.")

	log.Println("[RABBIT] All routes for RABBIT entity loaded successfully.")
}
