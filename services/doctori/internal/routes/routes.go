package routes

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/database"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/database/redis"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/middleware"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func SetupRoutes(parentCtx context.Context, dbConn database.Database, rdb *redis.RedisClient) *mux.Router {
	log.Println("[DOCTOR] Setting up rate limiter...")
	rateLimiter := middleware.NewRedisRateLimiter(rdb.GetClient(), parentCtx, utils.LIMITER_REQUESTS_ALLOWED, utils.LIMITER_MINUTE_MULTIPLIER*time.Minute)

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
	router.Handle(utils.CREATE_DOCTOR_ENDPOINT, middleware.ValidateDoctorInfo(doctorCreationHandler)).Methods("POST") // Creates a new doctor
	log.Println("[DOCTOR] Route POST", utils.CREATE_DOCTOR_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	doctorFetchAllHandler := http.HandlerFunc(doctorController.GetDoctors)
	router.HandleFunc(utils.FETCH_ALL_DOCTORS_ENDPOINT, doctorFetchAllHandler).Methods("GET") // Lists all doctors
	log.Println("[DOCTOR] Route GET", utils.FETCH_ALL_DOCTORS_ENDPOINT, "registered.")

	doctorFetchByIDHandler := http.HandlerFunc(doctorController.GetDoctorByID)
	router.HandleFunc(utils.FETCH_DOCTOR_BY_ID_ENDPOINT, doctorFetchByIDHandler).Methods("GET") // Lists all doctors
	log.Println("[DOCTOR] Route GET", utils.FETCH_DOCTOR_BY_ID_ENDPOINT, "registered.")

	doctorFetchByEmailHandler := http.HandlerFunc(doctorController.GetDoctorByEmail)
	router.Handle(utils.FETCH_DOCTOR_BY_EMAIL_ENDPOINT, middleware.ValidateEmail(doctorFetchByEmailHandler)).Methods("GET")
	log.Println("[DOCTOR] Route GET", utils.FETCH_DOCTOR_BY_EMAIL_ENDPOINT, "registered.")

	doctorFetchByUserIDHandler := http.HandlerFunc(doctorController.GetDoctorByUserID)
	router.HandleFunc(utils.FETCH_DOCTOR_BY_USER_ID_ENDPOINT, doctorFetchByUserIDHandler).Methods("GET")
	log.Println("[DOCTOR] Route GET", utils.FETCH_DOCTOR_BY_USER_ID_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	doctorUpdateByIDHandler := http.HandlerFunc(doctorController.UpdateDoctorByID)
	router.Handle(utils.UPDATE_DOCTOR_BY_ID_ENDPOINT, middleware.ValidateDoctorInfo(doctorUpdateByIDHandler)).Methods("PUT") // Updates a specific doctor
	log.Println("[DOCTOR] Route PUT", utils.UPDATE_DOCTOR_BY_ID_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	doctorDeleteByIDHandler := http.HandlerFunc(doctorController.DeleteDoctorByID)
	router.Handle(utils.DELETE_DOCTOR_BY_ID_ENDPOINT, doctorDeleteByIDHandler).Methods("DELETE") // Deletes a doctor
	log.Println("[DOCTOR] Route DELETE", utils.DELETE_DOCTOR_BY_ID_ENDPOINT, "registered.")

	log.Println("[DOCTOR] All CRUD routes for Doctor entity loaded successfully.")
}
