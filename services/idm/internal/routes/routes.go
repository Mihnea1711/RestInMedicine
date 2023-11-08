package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/idm/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/idm/internal/database"
	"github.com/mihnea1711/POS_Project/services/idm/internal/database/redis"
	"github.com/mihnea1711/POS_Project/services/idm/internal/middleware"
)

func SetupRoutes(dbConn database.Database, rdb *redis.RedisClient) *mux.Router {
	log.Println("[IDM] Setting up rate limiter...")
	rateLimiter := middleware.NewRedisRateLimiter(rdb, 10, time.Minute) // Here, I'm allowing 10 requests per minute.

	log.Println("[IDM] Setting up routes...")
	router := mux.NewRouter()
	router.Use(rateLimiter.Limit)
	router.Use(middleware.RouteLogger)

	idmController := &controllers.IDMController{
		DbConn:    dbConn,
		RedisConn: rdb,
	}

	loadRoutes(router, idmController)

	log.Println("[IDM] Routes setup completed.")
	return router
}

// loadRoutes loads all the CRUD routes for the IDM module
func loadRoutes(router *mux.Router, idmController *controllers.IDMController) {
	log.Println("[IDM] Loading CRUD routes for IDM module...")

	// User Registration
	userRegistrationHandler := http.HandlerFunc(idmController.RegisterUser)
	router.Handle("/idm/register", middleware.ValidateRegisterUserInfo(userRegistrationHandler)).Methods("POST")
	log.Println("[IDM] Route POST /idm/register registered.")

	// User Login
	userLoginHandler := http.HandlerFunc(idmController.LoginUser)
	router.Handle("/idm/login", userLoginHandler).Methods("POST")
	log.Println("[IDM] Route POST /idm/login registered.")

	// Get Users
	getUsersHandler := http.HandlerFunc(idmController.GetUsers)
	router.Handle("/idm/user", getUsersHandler).Methods("GET")
	log.Println("[IDM] Route GET /idm/user registered.")

	// Get User by ID
	getUserByIDHandler := http.HandlerFunc(idmController.GetUserByID)
	router.Handle("/idm/user/{id}", getUserByIDHandler).Methods("GET")
	log.Println("[IDM] Route GET /idm/user/{id} registered.")

	// Update User by ID
	updateUserByIDHandler := http.HandlerFunc(idmController.UpdateUserByID)
	router.Handle("/idm/user/{id}", updateUserByIDHandler).Methods("PUT")
	log.Println("[IDM] Route PUT /idm/user/{id} registered.")

	// Delete User by ID
	deleteUserByIDHandler := http.HandlerFunc(idmController.DeleteUserByID)
	router.Handle("/idm/user/{id}", deleteUserByIDHandler).Methods("DELETE")
	log.Println("[IDM] Route DELETE /idm/user/{id} registered.")

	// Get User Role
	getUserRoleHandler := http.HandlerFunc(idmController.GetUserRole)
	router.Handle("/idm/user/{id}/role", getUserRoleHandler).Methods("GET")
	log.Println("[IDM] Route GET /idm/user/{id}/role registered.")

	// Get User Token
	getUserTokenHandler := http.HandlerFunc(idmController.GetUserToken)
	router.Handle("/idm/user/{id}/token", getUserTokenHandler).Methods("GET")
	log.Println("[IDM] Route GET /idm/user/{id}/token registered.")

	// Add User to Blacklist
	addUserToBlacklistHandler := http.HandlerFunc(idmController.AddUserToBlacklist)
	router.Handle("/idm/blacklist/add", addUserToBlacklistHandler).Methods("POST")
	log.Println("[IDM] Route POST /idm/blacklist/add registered.")

	// Remove User from Blacklist
	removeUserFromBlacklistHandler := http.HandlerFunc(idmController.RemoveUserFromBlacklist)
	router.Handle("/idm/blacklist/remove", removeUserFromBlacklistHandler).Methods("POST")
	log.Println("[IDM] Route POST /idm/blacklist/remove registered.")

	// Change User Password
	changeUserPasswordHandler := http.HandlerFunc(idmController.UpdateUserPassword)
	router.Handle("/idm/user/{id}/password", changeUserPasswordHandler).Methods("PUT")
	log.Println("[IDM] Route PUT /idm/user/{id}/password registered.")

	// Change User Role
	changeUserRoleHandler := http.HandlerFunc(idmController.UpdateUserRole)
	router.Handle("/idm/user/{id}/role", changeUserRoleHandler).Methods("PUT")
	log.Println("[IDM] Route PUT /idm/user/{id}/role registered.")

	log.Println("[IDM] All CRUD routes for IDM module loaded successfully.")
}
