package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/database"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/middleware"
	"github.com/redis/go-redis/v9"
)

func SetupRoutes(dbConn database.Database, rdb *redis.Client) *mux.Router {
	log.Println("[PACIENT] Setting up rate limiter...")
	rateLimiter := middleware.NewRedisRateLimiter(rdb, 10, time.Minute) // allowing 10 requests per minute.

	log.Println("[PACIENT] Setting up routes...")
	router := mux.NewRouter()
	router.Use(rateLimiter.Limit)
	router.Use(middleware.RouteLogger)

	pacientController := &controllers.PacientController{
		DbConn: dbConn,
	}

	loadCrudRoutes(router, pacientController)

	log.Println("[PACIENT] Routes setup completed.")
	return router
}

// loadCrudRoutes loads all the CRUD routes for the pacient entity
func loadCrudRoutes(router *mux.Router, pacientController *controllers.PacientController) {
	log.Println("[PACIENT] Loading CRUD routes for Pacient entity...")

	// ---------------------------------------------------------- Create --------------------------------------------------------------
	pacientCreationHandler := http.HandlerFunc(pacientController.CreatePacient)
	router.Handle("/pacienti", middleware.ValidatePacientInfo(pacientCreationHandler)).Methods("POST") // Creates a new pacient
	log.Println("[PACIENT] Route POST /pacienti registered.")

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	pacientFetchAllHandler := http.HandlerFunc(pacientController.GetPacienti)
	router.HandleFunc("/pacienti", pacientFetchAllHandler).Methods("GET") // Lists all pacienti
	log.Println("[PACIENT] Route GET /pacienti registered.")

	pacientFetchByIDHandler := http.HandlerFunc(pacientController.GetPacientByID)
	router.HandleFunc("/pacienti/{id}", pacientFetchByIDHandler).Methods("GET") // Lists all pacienti
	log.Println("[PACIENT] Route GET /pacienti/{id} registered.")

	pacientFetchByEmailHandler := http.HandlerFunc(pacientController.GetPacientByEmail)
	router.Handle("/pacienti/email/{email}", middleware.ValidateEmail(pacientFetchByEmailHandler)).Methods("GET")
	log.Println("[PACIENT] Route GET /pacienti/email/{email} registered.")

	pacientFetchByUserIDHandler := http.HandlerFunc(pacientController.GetPacientByUserID)
	router.HandleFunc("/pacienti/users/{id}", pacientFetchByUserIDHandler).Methods("GET")
	log.Println("[PACIENT] Route GET /pacienti/users/{id} registered.")

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	pacientUpdateByIDHandler := http.HandlerFunc(pacientController.UpdatePacientByID)
	router.Handle("/pacienti/{id}", middleware.ValidatePacientInfo(pacientUpdateByIDHandler)).Methods("PUT") // Updates a specific pacient
	log.Println("[PACIENT] Route PUT /pacienti/{id} registered.")

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	pacientDeleteByIDHandler := http.HandlerFunc(pacientController.DeletePacientByID)
	router.Handle("/pacienti/{id}", pacientDeleteByIDHandler).Methods("DELETE") // Deletes a pacient
	log.Println("[PACIENT] Route DELETE /pacienti/{id} registered.")

	log.Println("[PACIENT] All CRUD routes for Pacient entity loaded successfully.")
}
