package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/middleware/authorization"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/middleware/validation"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/config"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// loadUserRoutes loads all the CRUD routes for the User entity
func loadUserRoutes(router *mux.Router, gatewayController *controllers.GatewayController, jwtConfig config.JWTConfig) {
	// ---------------------------------------------------------- Create --------------------------------------------------------------
	registerUserHandler := http.HandlerFunc(gatewayController.RegisterUser)
	router.Handle(utils.REGISTER_USER_ENDPOINT, validation.ValidateRegistrationData(registerUserHandler)).Methods("POST")
	log.Printf("[GATEWAY] Route POST %s registered.\n", utils.REGISTER_USER_ENDPOINT)

	loginUserHandler := http.HandlerFunc(gatewayController.LoginUser)
	router.Handle(utils.LOGIN_USER_ENDPOINT, validation.ValidateLoginData(loginUserHandler)).Methods("POST")
	log.Printf("[GATEWAY] Route POST %s registered.\n", utils.LOGIN_USER_ENDPOINT)

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	getAllUserHandler := http.HandlerFunc(gatewayController.GetAllUsers)
	router.Handle(utils.GET_ALL_USERS_ENDPOINT, authorization.AdminOnlyMiddleware(jwtConfig, getAllUserHandler)).Methods("GET")
	log.Printf("[GATEWAY] Route GET %s registered.\n", utils.GET_ALL_USERS_ENDPOINT)

	getByIDUserHandler := http.HandlerFunc(gatewayController.GetUserByID)
	router.Handle(utils.GET_USER_BY_ID_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, getByIDUserHandler)).Methods("GET")
	log.Printf("[GATEWAY] Route GET %s registered.\n", utils.GET_USER_BY_ID_ENDPOINT)

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	updateUserHandler := http.HandlerFunc(gatewayController.UpdateUser)
	router.Handle(utils.UPDATE_USER_BY_ID_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, validation.ValidateUserData(updateUserHandler))).Methods("PUT")
	log.Printf("[GATEWAY] Route PUT %s registered.\n", utils.UPDATE_USER_BY_ID_ENDPOINT)

	updatePasswordHandler := http.HandlerFunc(gatewayController.UpdatePassword)
	router.Handle(utils.UPDATE_PASSWORD_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, validation.ValidatePasswordData(updatePasswordHandler))).Methods("POST")
	log.Printf("[GATEWAY] Route POST %s registered.\n", utils.UPDATE_PASSWORD_ENDPOINT)

	updateRoleHandler := http.HandlerFunc(gatewayController.UpdateRole)
	router.Handle(utils.UPDATE_ROLE_ENDPOINT, authorization.AdminOnlyMiddleware(jwtConfig, validation.ValidateRoleData(updateRoleHandler))).Methods("POST")
	log.Printf("[GATEWAY] Route POST %s registered.\n", utils.UPDATE_ROLE_ENDPOINT)

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	deleteUserHandler := http.HandlerFunc(gatewayController.DeleteUser)
	router.Handle(utils.DELETE_USER_BY_ID_ENDPOINT, authorization.AdminOnlyMiddleware(jwtConfig, deleteUserHandler)).Methods("DELETE")
	log.Printf("[GATEWAY] Route DELETE %s registered.\n", utils.DELETE_USER_BY_ID_ENDPOINT)

	// ---------------------------------------------------------- Blacklist --------------------------------------------------------------
	addToBlacklistHandler := http.HandlerFunc(gatewayController.AddToBlacklist)
	router.Handle(utils.ADD_TO_BLACKLIST_ENDPOINT, authorization.AdminOnlyMiddleware(jwtConfig, validation.ValidateBlacklistData(addToBlacklistHandler))).Methods("POST")
	log.Printf("[GATEWAY] Route POST %s registered.\n", utils.ADD_TO_BLACKLIST_ENDPOINT)

	checkBlacklistHandler := http.HandlerFunc(gatewayController.CheckBlacklist)
	router.Handle(utils.CHECK_BLACKLIST_ENDPOINT, authorization.AdminOnlyMiddleware(jwtConfig, checkBlacklistHandler)).Methods("GET")
	log.Printf("[GATEWAY] Route POST %s registered.\n", utils.CHECK_BLACKLIST_ENDPOINT)

	removeFromBlacklistHandler := http.HandlerFunc(gatewayController.RemoveFromBlacklist)
	router.Handle(utils.DELETE_FROM_BLACKLIST_ENDPOINT, authorization.AdminOnlyMiddleware(jwtConfig, removeFromBlacklistHandler)).Methods("DELETE")
	log.Printf("[GATEWAY] Route POST %s registered.\n", utils.DELETE_FROM_BLACKLIST_ENDPOINT)
}
