package routes

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/consultatii/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/consultatii/internal/database"
	"github.com/mihnea1711/POS_Project/services/consultatii/internal/database/redis"
	"github.com/mihnea1711/POS_Project/services/consultatii/internal/middleware"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
)

func SetupRoutes(ctx context.Context, dbConn database.Database, rdb *redis.RedisClient) *mux.Router {
	log.Println("[CONSULTATION] Setting up rate limiter...")
	rateLimiter := middleware.NewRedisRateLimiter(ctx, rdb, utils.REQUEST_RATE, utils.REQUEST_WINDOW_DURATION_MULTIPLIER*time.Minute)
	log.Println("[CONSULTATION] Rate limiter set up successfully.")

	log.Println("[CONSULTATION] Setting up routes...")
	router := mux.NewRouter()

	router.Use(rateLimiter.Limit)
	log.Println("[CONSULTATION] Rate limiter middleware set up successfully.")

	router.Use(middleware.RouteLogger)
	log.Println("[CONSULTATION] Route logger middleware set up successfully.")

	router.Use(middleware.SanitizeInputMiddleware) // comment this out if you want to see pretty JSON :)
	log.Println("[CONSULTATION] Input sanitizer middleware set up successfully.")

	consultatieController := &controllers.ConsultationController{
		DbConn: dbConn,
	}

	loadCrudRoutes(router, consultatieController)

	log.Println("[CONSULTATION] Routes setup completed.")
	return router
}

// loadCrudRoutes loads all the CRUD routes for the consultatie entity
func loadCrudRoutes(router *mux.Router, consultatieController *controllers.ConsultationController) {
	log.Println("[CONSULTATION] Loading CRUD routes for consultatie entity...")

	// ---------------------------------------------------------- Health --------------------------------------------------------------
	healthHandler := http.HandlerFunc(consultatieController.HealthCheck)
	router.Handle(utils.HEALTH_CHECK_ENDPOINT, healthHandler).Methods("GET")
	log.Printf("[CONSULTATION] Route DELETE %s registered.", utils.HEALTH_CHECK_ENDPOINT)

	// ---------------------------------------------------------- Create --------------------------------------------------------------
	consultatieCreationHandler := http.HandlerFunc(consultatieController.CreateConsultation)
	router.Handle(utils.INSERT_CONSULTATIE_ENDPOINT, middleware.ValidateConsultationInfo(consultatieCreationHandler)).Methods("POST")
	log.Print("[CONSULTATION] Route POST /consultatii registered.")

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	consultatieFetchAllHandler := http.HandlerFunc(consultatieController.GetConsultations)
	router.HandleFunc(utils.FETCH_ALL_CONSULTATII_ENDPOINT, consultatieFetchAllHandler).Methods("GET")
	log.Printf("[CONSULTATION] Route GET %s registered.", utils.FETCH_ALL_CONSULTATII_ENDPOINT)

	consultatieFetchByIDHandler := http.HandlerFunc(consultatieController.GetConsultationByID)
	router.HandleFunc(utils.FETCH_CONSULTATIE_BY_ID_ENDPOINT, consultatieFetchByIDHandler).Methods("GET")
	log.Printf("[CONSULTATION] Route GET %s registered.", utils.FETCH_CONSULTATIE_BY_ID_ENDPOINT)

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	consultatieUpdateByIDHandler := http.HandlerFunc(consultatieController.UpdateConsultationByID)
	router.Handle(utils.UPDATE_CONSULTATIE_BY_ID_ENDPOINT, middleware.ValidateConsultationInfo(consultatieUpdateByIDHandler)).Methods("PUT")
	log.Printf("[CONSULTATION] Route PUT %s registered.", utils.UPDATE_CONSULTATIE_BY_ID_ENDPOINT)

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	consultatieDeleteByIDHandler := http.HandlerFunc(consultatieController.DeleteConsultationByID)
	router.Handle(utils.DELETE_CONSULTATIE_BY_ID_ENDPOINT, consultatieDeleteByIDHandler).Methods("DELETE")
	log.Printf("[CONSULTATION] Route DELETE %s registered.", utils.DELETE_CONSULTATIE_BY_ID_ENDPOINT)

	consultatieDeleteByPatientOrDoctorIDHandler := http.HandlerFunc(consultatieController.DeleteConsultationByPatientOrDoctorID)
	router.Handle(utils.DELETE_CONSULTATIE_BY_PATIENT_DOCTOR_ID_ENDPOINT, consultatieDeleteByPatientOrDoctorIDHandler).Methods("DELETE")
	log.Printf("[CONSULTATION] Route DELETE %s registered.", utils.DELETE_CONSULTATIE_BY_PATIENT_DOCTOR_ID_ENDPOINT)

	log.Println("[CONSULTATION] All CRUD routes for consultatie entity loaded successfully.")
}
