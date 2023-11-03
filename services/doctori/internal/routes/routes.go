package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/database"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/middleware"
	"github.com/redis/go-redis/v9"
)

func SetupRoutes(dbConn database.Database, rdb *redis.Client) *mux.Router {
	log.Println("[DOCTOR] Setting up rate limiter...")
	rateLimiter := middleware.NewRedisRateLimiter(rdb, 10, time.Minute) // Here, I'm allowing 10 requests per minute.

	log.Println("[DOCTOR] Setting up routes...")
	router := mux.NewRouter()
	router.Use(rateLimiter.Limit)
	router.Use(middleware.RouteLogger)

	doctorController := &controllers.DoctorController{
		DbConn: dbConn,
	}

	loadCrudRoutes(router, doctorController)

	log.Println("[DOCTOR] Routes setup completed.")
	return router
}

// loadCrudRoutes loads all the CRUD routes for the Doctor entity
func loadCrudRoutes(router *mux.Router, doctorController *controllers.DoctorController) {
	log.Println("[DOCTOR] Loading CRUD routes for Doctor entity...")

	// ---------------------------------------------------------- Create --------------------------------------------------------------
	doctorCreationHandler := http.HandlerFunc(doctorController.CreateDoctor)
	router.Handle("/doctori", middleware.ValidateDoctorInfo(doctorCreationHandler)).Methods("POST") // Creates a new doctor
	log.Println("[DOCTOR] Route POST /doctori registered.")

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	doctorFetchAllHandler := http.HandlerFunc(doctorController.GetDoctors)
	router.HandleFunc("/doctori", doctorFetchAllHandler).Methods("GET") // Lists all doctors
	log.Println("[DOCTOR] Route GET /doctori registered.")

	doctorFetchByIDHandler := http.HandlerFunc(doctorController.GetDoctorByID)
	router.HandleFunc("/doctori/{id}", doctorFetchByIDHandler).Methods("GET") // Lists all doctors
	log.Println("[DOCTOR] Route GET /doctori/{id} registered.")

	doctorFetchByEmailHandler := http.HandlerFunc(doctorController.GetDoctorByEmail)
	router.Handle("/doctori/email/{email}", middleware.ValidateEmail(doctorFetchByEmailHandler)).Methods("GET")
	log.Println("[DOCTOR] Route GET /doctori/email/{email} registered.")

	doctorFetchByUserIDHandler := http.HandlerFunc(doctorController.GetDoctorByUserID)
	router.HandleFunc("/doctori/users/{id}", doctorFetchByUserIDHandler).Methods("GET")
	log.Println("[DOCTOR] Route GET /doctori/users/{id} registered.")

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	doctorUpdateByIDHandler := http.HandlerFunc(doctorController.UpdateDoctorByID)
	router.Handle("/doctori/{id}", middleware.ValidateDoctorInfo(doctorUpdateByIDHandler)).Methods("PUT") // Updates a specific doctor
	log.Println("[DOCTOR] Route PUT /doctori/{id} registered.")

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	doctorDeleteByIDHandler := http.HandlerFunc(doctorController.DeleteDoctorByID)
	router.Handle("/doctori/{id}", doctorDeleteByIDHandler).Methods("DELETE") // Deletes a doctor
	log.Println("[DOCTOR] Route DELETE /doctori/{id} registered.")

	log.Println("[DOCTOR] All CRUD routes for Doctor entity loaded successfully.")
}
