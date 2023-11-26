package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/middleware"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/utils"
)

func SetupRoutes() *mux.Router {
	log.Println("[RABBIT] Setting up routes...")
	router := mux.NewRouter()
	router.Use(middleware.RouteLogger)

	rabbitController := &controllers.RabbitController{}

	loadRoutes(router, rabbitController)

	log.Println("[RABBIT] Routes setup completed.")
	return router
}

func loadRoutes(router *mux.Router, rabbitController *controllers.RabbitController) {
	log.Println("[RABBIT] Loading routes for RABBIT entity...")

	// ---------------------------------------------------------- Create --------------------------------------------------------------
	healthCheckHandler := http.HandlerFunc(rabbitController.HandleHealthCheck)
	router.Handle(utils.HEALTH_CHECK_ENDPOINT, healthCheckHandler).Methods("GET") // Handles health check
	log.Println("[RABBIT] Route POST", utils.HEALTH_CHECK_ENDPOINT, "registered.")

	log.Println("[RABBIT] All routes for RABBIT entity loaded successfully.")
}
