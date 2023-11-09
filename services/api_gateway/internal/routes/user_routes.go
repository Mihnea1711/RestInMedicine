package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/controllers"
)

// loadRoutes loads all the CRUD routes for the GATEWAY entity
func loadUserRoutes(router *mux.Router, gatewayController *controllers.GatewayController) {

	// RegisterUser handles user registration.
	registerUserHandler := http.HandlerFunc(gatewayController.RegisterUser)
	router.Handle("/api/register", registerUserHandler).Methods("POST").Queries("role", "{role}")
	log.Println("[GATEWAY] Route POST /api/register registered.")

	// LoginUser handles user login.
	loginUserHandler := http.HandlerFunc(gatewayController.LoginUser)
	router.Handle("/api/login", loginUserHandler).Methods("POST")
	log.Println("[GATEWAY] Route POST /api/login registered.")

}
