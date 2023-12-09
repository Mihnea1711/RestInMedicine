package routes

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/programari/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/programari/internal/database"
	"github.com/mihnea1711/POS_Project/services/programari/internal/database/redis"
	"github.com/mihnea1711/POS_Project/services/programari/internal/middleware"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

func SetupRoutes(ctx context.Context, dbConn database.Database, rdb *redis.RedisClient) *mux.Router {
	log.Println("[APPOINTMENT] Setting up rate limiter...")
	rateLimiter := middleware.NewRedisRateLimiter(ctx, rdb.GetClient(), utils.LIMITER_REQUESTS_ALLOWED, utils.LIMITER_MINUTE_MULTIPLIER*time.Minute)
	log.Println("[APPOINTMENT] Rate limiter set up successfully.")

	log.Println("[APPOINTMENT] Setting up routes...")
	router := mux.NewRouter()

	router.Use(rateLimiter.Limit)
	log.Println("[APPOINTMENT] Rate limiter middleware set up successfully.")

	router.Use(middleware.RouteLogger)
	log.Println("[APPOINTMENT] Route logger middleware set up successfully.")

	router.Use(middleware.SanitizeInputMiddleware) // comment this out if you want to see pretty JSON :)
	log.Println("[APPOINTMENT] Input sanitizer middleware set up successfully.")

	appointmentsController := &controllers.AppointmentController{
		DbConn: dbConn,
	}

	loadCrudRoutes(router, appointmentsController)

	log.Println("[APPOINTMENT] Routes setup completed.")
	return router
}

// loadCrudRoutes loads all the CRUD routes for the Doctor entity
func loadCrudRoutes(router *mux.Router, appointmentController *controllers.AppointmentController) {
	log.Println("[APPOINTMENT] Loading CRUD routes for Appointment entity...")

	// ---------------------------------------------------------- Health --------------------------------------------------------------

	healthCheckHandler := http.HandlerFunc(appointmentController.HealthCheck)
	router.Handle(utils.HEALTH_CHECK_ENDPOINT, healthCheckHandler).Methods("GET")
	log.Println("[APPOINTMENT] Route GET", utils.HEALTH_CHECK_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Create --------------------------------------------------------------
	appointmentCreationHandler := http.HandlerFunc(appointmentController.CreateAppointment)
	router.Handle(utils.CREATE_APPOINTMENT_ENDPOINT, middleware.ValidateAppointmentInfo(appointmentCreationHandler)).Methods("POST") // Creates a new appointment
	log.Println("[APPOINTMENT] Route POST", utils.CREATE_APPOINTMENT_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	// // implement pagination for these
	// appointmentFetchAllHandler := http.HandlerFunc(appointmentController.GetAppointments)
	// router.HandleFunc(utils.FETCH_ALL_APPOINTMENTS_ENDPOINT, appointmentFetchAllHandler).Methods("GET") // Lists all appointments
	// log.Println("[APPOINTMENT] Route GET", utils.FETCH_ALL_APPOINTMENTS_ENDPOINT, "registered.")

	appointmentFetchAllHandler := http.HandlerFunc(appointmentController.GetAppointmentsByFilter)
	router.HandleFunc(utils.FETCH_ALL_APPOINTMENTS_ENDPOINT, appointmentFetchAllHandler).Methods("GET") // Lists all appointments
	log.Println("[APPOINTMENT] Route GET", utils.FETCH_ALL_APPOINTMENTS_ENDPOINT, "registered.")

	appointmentFetchByIDHandler := http.HandlerFunc(appointmentController.GetAppointmentByID)
	router.HandleFunc(utils.FETCH_APPOINTMENT_BY_ID_ENDPOINT, appointmentFetchByIDHandler).Methods("GET") // Get a specific appointment by ID
	log.Println("[APPOINTMENT] Route GET", utils.FETCH_APPOINTMENT_BY_ID_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	appointmentUpdateByIDHandler := http.HandlerFunc(appointmentController.UpdateAppointmentByID)
	router.Handle(utils.UPDATE_APPOINTMENT_BY_ID_ENDPOINT, middleware.ValidateAppointmentInfo(appointmentUpdateByIDHandler)).Methods("PUT") // Updates a specific appointment
	log.Println("[APPOINTMENT] Route PUT", utils.UPDATE_APPOINTMENT_BY_ID_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	appointmentDeleteByIDHandler := http.HandlerFunc(appointmentController.DeleteAppointmentByID)
	router.Handle(utils.DELETE_APPOINTMENT_BY_ID_ENDPOINT, appointmentDeleteByIDHandler).Methods("DELETE") // Deletes a appointment
	log.Println("[APPOINTMENT] Route DELETE", utils.DELETE_APPOINTMENT_BY_ID_ENDPOINT, "registered.")

	log.Println("[APPOINTMENT] All CRUD routes for Appointment entity loaded successfully.")
}
