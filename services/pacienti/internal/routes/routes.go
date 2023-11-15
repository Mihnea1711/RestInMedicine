package routes

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/database"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/database/redis"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/middleware"
)

func SetupRoutes(dbConn database.Database, rdb *redis.RedisClient, parentCtx context.Context) *mux.Router {
	log.Println("[PATIENT] Setting up rate limiter...")
	rateLimiter := middleware.NewRedisRateLimiter(rdb.GetClient(), parentCtx, 10, time.Minute) // allowing 10 requests per minute.

	log.Println("[PATIENT] Setting up routes...")
	router := mux.NewRouter()
	router.Use(rateLimiter.Limit)
	router.Use(middleware.RouteLogger)

	pacientController := &controllers.PacientController{
		DbConn: dbConn,
	}

	loadCrudRoutes(router, pacientController)

	log.Println("[PATIENT] Routes setup completed.")
	return router
}

// loadCrudRoutes loads all the CRUD routes for the pacient entity
func loadCrudRoutes(router *mux.Router, pacientController *controllers.PacientController) {
	log.Println("[PATIENT] Loading CRUD routes for Pacient entity...")

	// ---------------------------------------------------------- Create --------------------------------------------------------------
	pacientCreationHandler := http.HandlerFunc(pacientController.CreatePacient)
	router.Handle("/patients", middleware.ValidatePacientInfo(pacientCreationHandler)).Methods("POST") // Creates a new pacient
	log.Println("[PATIENT] Route POST /patients registered.")

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	pacientFetchAllHandler := http.HandlerFunc(pacientController.GetPacienti)
	router.HandleFunc("/patients", pacientFetchAllHandler).Methods("GET") // Lists all patients
	log.Println("[PATIENT] Route GET /patients registered.")

	pacientFetchByIDHandler := http.HandlerFunc(pacientController.GetPacientByID)
	router.HandleFunc("/patients/{id}", pacientFetchByIDHandler).Methods("GET") // Lists all patients
	log.Println("[PATIENT] Route GET /patients/{id} registered.")

	pacientFetchByEmailHandler := http.HandlerFunc(pacientController.GetPacientByEmail)
	router.Handle("/patients/email/{email}", middleware.ValidateEmail(pacientFetchByEmailHandler)).Methods("GET")
	log.Println("[PATIENT] Route GET /patients/email/{email} registered.")

	pacientFetchByUserIDHandler := http.HandlerFunc(pacientController.GetPacientByUserID)
	router.HandleFunc("/patients/users/{id}", pacientFetchByUserIDHandler).Methods("GET")
	log.Println("[PATIENT] Route GET /patients/users/{id} registered.")

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	pacientUpdateByIDHandler := http.HandlerFunc(pacientController.UpdatePacientByID)
	router.Handle("/patients/{id}", middleware.ValidatePacientInfo(pacientUpdateByIDHandler)).Methods("PUT") // Updates a specific pacient
	log.Println("[PATIENT] Route PUT /patients/{id} registered.")

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	pacientDeleteByIDHandler := http.HandlerFunc(pacientController.DeletePacientByID)
	router.Handle("/patients/{id}", pacientDeleteByIDHandler).Methods("DELETE") // Deletes a pacient
	log.Println("[PATIENT] Route DELETE /patients/{id} registered.")

	log.Println("[PATIENT] All CRUD routes for Patient entity loaded successfully.")
}
