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

	loadCrudRoutes(router, gatewayController)

	log.Println("[GATEWAY] Routes setup completed.")
	return router
}

// loadCrudRoutes loads all the CRUD routes for the GATEWAY entity
func loadCrudRoutes(router *mux.Router, gatewayController *controllers.GatewayController) {
	log.Println("[GATEWAY] Loading CRUD routes for GATEWAY entity...")

	log.Println("[GATEWAY] All CRUD routes for GATEWAY entity loaded successfully.")
}
