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
	router.Handle("/api/users", registerUserHandler).Methods("POST")
	log.Println("[GATEWAY] Route POST /api/users registered.")

	// LoginUser handles user login.
	loginUserHandler := http.HandlerFunc(gatewayController.LoginUser)
	router.Handle("/api/login", loginUserHandler).Methods("POST")
	log.Println("[GATEWAY] Route POST /api/login registered.")

	// GetAllUsers handles fetching all users.
	getAllUserHandler := http.HandlerFunc(gatewayController.GetAllUsers)
	router.Handle("/api/users", getAllUserHandler).Methods("GET")
	log.Println("[GATEWAY] Route GET /api/users registered.")

	// GetByIDUser handles fetching a user by ID.
	getByIDUserHandler := http.HandlerFunc(gatewayController.GetUserByID)
	router.Handle("/api/users/{userID}", getByIDUserHandler).Methods("GET")
	log.Println("[GATEWAY] Route GET /api/users/{userID} registered.")

	// UpdateUser handles updating a user.
	updateUserHandler := http.HandlerFunc(gatewayController.UpdateUser)
	router.Handle("/api/users/{userID}", updateUserHandler).Methods("PUT")
	log.Println("[GATEWAY] Route PUT /api/users/{userID} registered.")

	// DeleteUser handles deleting a user.
	deleteUserHandler := http.HandlerFunc(gatewayController.DeleteUser)
	router.Handle("/api/users/{userID}", deleteUserHandler).Methods("DELETE")
	log.Println("[GATEWAY] Route DELETE /api/users/{userID} registered.")

	// UpdatePassword handles updating a user's password.
	updatePasswordHandler := http.HandlerFunc(gatewayController.UpdatePassword)
	router.Handle("/api/users/{userID}/update-password", updatePasswordHandler).Methods("POST")
	log.Println("[GATEWAY] Route POST /api/users/{userID}/update-password registered.")

	// UpdateRole handles updating a user's role.
	updateRoleHandler := http.HandlerFunc(gatewayController.UpdateRole)
	router.Handle("/api/users/{userID}/update-role", updateRoleHandler).Methods("POST")
	log.Println("[GATEWAY] Route POST /api/users/{userID}/update-password registered.")
}
