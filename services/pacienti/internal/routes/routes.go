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
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func SetupRoutes(parentCtx context.Context, dbConn database.Database, rdb *redis.RedisClient) *mux.Router {
	log.Println("[PATIENT] Setting up rate limiter...")
	rateLimiter := middleware.NewRedisRateLimiter(rdb.GetClient(), parentCtx, utils.LIMITER_REQUESTS_ALLOWED, utils.LIMITER_MINUTE_MULTIPLIER*time.Minute)
	log.Println("[PATIENT] Rate limiter set up successfully.")

	log.Println("[PATIENT] Setting up routes...")
	router := mux.NewRouter()

	router.Use(rateLimiter.Limit)
	log.Println("[PATIENT] Rate limiter middleware set up successfully.")

	router.Use(middleware.RouteLogger)
	log.Println("[PATIENT] Route logger middleware set up successfully.")

	router.Use(middleware.SanitizeInputMiddleware) // comment this out if you want to see pretty JSON :)
	log.Println("[PATIENT] Input sanitizer middleware set up successfully.")

	pacientController := &controllers.PatientController{
		DbConn: dbConn,
	}

	loadCrudRoutes(router, pacientController)

	log.Println("[PATIENT] Routes loaded successfully.")
	return router
}

// loadCrudRoutes loads all the CRUD routes for the pacient entity
func loadCrudRoutes(router *mux.Router, pacientController *controllers.PatientController) {
	log.Println("[PATIENT] Loading CRUD routes for Pacient entity...")

	pacientCreationHandler := http.HandlerFunc(pacientController.CreatePatient)
	router.Handle(utils.CREATE_PATIENT_ENDPOINT, middleware.ValidatePacientInfo(pacientCreationHandler)).Methods("POST")
	log.Println("[PATIENT] Route POST", utils.CREATE_PATIENT_ENDPOINT, "registered.")

	pacientFetchAllHandler := http.HandlerFunc(pacientController.GetPatients)
	router.HandleFunc(utils.FETCH_ALL_PATIENTS_ENDPOINT, pacientFetchAllHandler).Methods("GET")
	log.Println("[PATIENT] Route GET", utils.FETCH_ALL_PATIENTS_ENDPOINT, "registered.")

	pacientFetchByEmailHandler := http.HandlerFunc(pacientController.GetPatientByEmail)
	router.Handle(utils.FETCH_PATIENT_BY_EMAIL_ENDPOINT, middleware.ValidateEmail(pacientFetchByEmailHandler)).Methods("GET")
	log.Println("[PATIENT] Route GET", utils.FETCH_PATIENT_BY_EMAIL_ENDPOINT, "registered.")

	pacientFetchByUserIDHandler := http.HandlerFunc(pacientController.GetPatientByUserID)
	router.HandleFunc(utils.FETCH_PATIENT_BY_USER_ID_ENDPOINT, pacientFetchByUserIDHandler).Methods("GET")
	log.Println("[PATIENT] Route GET", utils.FETCH_PATIENT_BY_USER_ID_ENDPOINT, "registered.")

	healthCheckHandler := http.HandlerFunc(pacientController.HealthCheck)
	router.Handle(utils.HEALTH_CHECK_ENDPOINT, healthCheckHandler).Methods("GET")
	log.Println("[PATIENT] Route GET", utils.HEALTH_CHECK_ENDPOINT, "registered.")

	pacientFetchByIDHandler := http.HandlerFunc(pacientController.GetPatientByID)
	router.HandleFunc(utils.FETCH_PATIENT_BY_ID_ENDPOINT, pacientFetchByIDHandler).Methods("GET")
	log.Println("[PATIENT] Route GET", utils.FETCH_PATIENT_BY_ID_ENDPOINT, "registered.")

	pacientUpdateByIDHandler := http.HandlerFunc(pacientController.UpdatePatientByID)
	router.Handle(utils.UPDATE_PATIENT_BY_ID_ENDPOINT, middleware.ValidatePacientInfo(pacientUpdateByIDHandler)).Methods("PUT")
	log.Println("[PATIENT] Route PUT", utils.UPDATE_PATIENT_BY_ID_ENDPOINT, "registered.")

	pacientDeleteByIDHandler := http.HandlerFunc(pacientController.DeletePatientByID)
	router.Handle(utils.DELETE_PATIENT_BY_ID_ENDPOINT, pacientDeleteByIDHandler).Methods("DELETE")
	log.Println("[PATIENT] Route DELETE", utils.DELETE_PATIENT_BY_ID_ENDPOINT, "registered.")

	pacientDeleteByUserIDHandler := http.HandlerFunc(pacientController.DeletePatientByUserID)
	router.Handle(utils.DELETE_PATIENT_BY_USER_ID_ENDPOINT, pacientDeleteByUserIDHandler).Methods("DELETE")
	log.Println("[PATIENT] Route DELETE", utils.DELETE_PATIENT_BY_USER_ID_ENDPOINT, "registered.")

	log.Println("[PATIENT] All CRUD routes for Patient entity loaded successfully.")
}
