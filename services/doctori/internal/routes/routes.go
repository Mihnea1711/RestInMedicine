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
	log.Println("[DOCTOR] Rate limiter set up successfully.")

	log.Println("[DOCTOR] Setting up routes...")
	router := mux.NewRouter()

	log.Println("[DOCTOR] Rate limiter middleware set up successfully.")
	router.Use(rateLimiter.Limit)

	log.Println("[DOCTOR] Route logger middleware set up successfully.")
	router.Use(middleware.RouteLogger)

	log.Println("[DOCTOR] Input sanitizer middleware set up successfully.")
	router.Use(middleware.SanitizeInputMiddleware) // comment this out if you want to see pretty JSON :)

	doctorController := &controllers.DoctorController{
		DbConn: dbConn,
	}

	loadCrudRoutes(router, doctorController)

	log.Println("[DOCTOR] Routes loaded successfully.")
	return router
}

// loadCrudRoutes loads all the CRUD routes for the Doctor entity
func loadCrudRoutes(router *mux.Router, doctorController *controllers.DoctorController) {
	log.Println("[DOCTOR] Loading CRUD routes for Doctor entity...")

	doctorCreationHandler := http.HandlerFunc(doctorController.CreateDoctor)
	router.Handle(utils.CREATE_DOCTOR_ENDPOINT, middleware.ValidateDoctorInfo(doctorCreationHandler)).Methods("POST")
	log.Println("[DOCTOR] Route POST", utils.CREATE_DOCTOR_ENDPOINT, "registered.")

	doctorFetchAllHandler := http.HandlerFunc(doctorController.GetDoctors)
	router.HandleFunc(utils.FETCH_ALL_DOCTORS_ENDPOINT, doctorFetchAllHandler).Methods("GET")
	log.Println("[DOCTOR] Route GET", utils.FETCH_ALL_DOCTORS_ENDPOINT, "registered.")

	doctorFetchByEmailHandler := http.HandlerFunc(doctorController.GetDoctorByEmail)
	router.Handle(utils.FETCH_DOCTOR_BY_EMAIL_ENDPOINT, middleware.ValidateEmail(doctorFetchByEmailHandler)).Methods("GET")
	log.Println("[DOCTOR] Route GET", utils.FETCH_DOCTOR_BY_EMAIL_ENDPOINT, "registered.")

	doctorFetchByUserIDHandler := http.HandlerFunc(doctorController.GetDoctorByUserID)
	router.HandleFunc(utils.FETCH_DOCTOR_BY_USER_ID_ENDPOINT, doctorFetchByUserIDHandler).Methods("GET")
	log.Println("[DOCTOR] Route GET", utils.FETCH_DOCTOR_BY_USER_ID_ENDPOINT, "registered.")

	healthCheckHandler := http.HandlerFunc(doctorController.HealthCheck)
	router.Handle(utils.HEALTH_CHECK_ENDPOINT, healthCheckHandler).Methods("GET")
	log.Println("[DOCTOR] Route GET", utils.HEALTH_CHECK_ENDPOINT, "registered.")

	doctorFetchByIDHandler := http.HandlerFunc(doctorController.GetDoctorByID)
	router.HandleFunc(utils.FETCH_DOCTOR_BY_ID_ENDPOINT, doctorFetchByIDHandler).Methods("GET")
	log.Println("[DOCTOR] Route GET", utils.FETCH_DOCTOR_BY_ID_ENDPOINT, "registered.")

	doctorUpdateByIDHandler := http.HandlerFunc(doctorController.UpdateDoctorByID)
	router.Handle(utils.UPDATE_DOCTOR_BY_ID_ENDPOINT, middleware.ValidateDoctorInfo(doctorUpdateByIDHandler)).Methods("PUT")
	log.Println("[DOCTOR] Route PUT", utils.UPDATE_DOCTOR_BY_ID_ENDPOINT, "registered.")

	toggleActivityHandler := http.HandlerFunc(doctorController.ToggleDoctorActivity)
	router.Handle(utils.TOGGLE_DOCTOR_ACTIVITY_ENDPOINT, middleware.ValidateDoctorActivityInfo(toggleActivityHandler)).Methods("PATCH")
	log.Println("[PATIENT] Route POST", utils.TOGGLE_DOCTOR_ACTIVITY_ENDPOINT, "registered.")

	doctorDeleteByUserIDHandler := http.HandlerFunc(doctorController.DeleteDoctorByUserID)
	router.Handle(utils.DELETE_DOCTOR_BY_USER_ID_ENDPOINT, doctorDeleteByUserIDHandler).Methods("DELETE")
	log.Println("[DOCTOR] Route DELETE", utils.DELETE_DOCTOR_BY_USER_ID_ENDPOINT, "registered.")

	doctorDeleteByIDHandler := http.HandlerFunc(doctorController.DeleteDoctorByID)
	router.Handle(utils.DELETE_DOCTOR_BY_ID_ENDPOINT, doctorDeleteByIDHandler).Methods("DELETE")
	log.Println("[DOCTOR] Route DELETE", utils.DELETE_DOCTOR_BY_ID_ENDPOINT, "registered.")

	log.Println("[DOCTOR] All CRUD routes for Doctor entity loaded successfully.")
}
