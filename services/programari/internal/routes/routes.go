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
	log.Println("[APPOINTMENT] Setting up routes...")
	router := mux.NewRouter()
	router.Use(rateLimiter.Limit)
	router.Use(middleware.RouteLogger)

	programariController := &controllers.ProgramareController{
		DbConn: dbConn,
	}

	loadCrudRoutes(router, programariController)

	log.Println("[APPOINTMENT] Routes setup completed.")
	return router
}

// loadCrudRoutes loads all the CRUD routes for the Doctor entity
func loadCrudRoutes(router *mux.Router, programareController *controllers.ProgramareController) {
	log.Println("[APPOINTMENT] Loading CRUD routes for Appointment entity...")

	// ---------------------------------------------------------- Create --------------------------------------------------------------
	programareCreationHandler := http.HandlerFunc(programareController.CreateProgramare)
	router.Handle(utils.CREATE_APPOINTMENT_ENDPOINT, middleware.ValidateProgramareInfo(programareCreationHandler)).Methods("POST") // Creates a new programare
	log.Println("[APPOINTMENT] Route POST", utils.CREATE_APPOINTMENT_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	// implement pagination for these
	programareFetchAllHandler := http.HandlerFunc(programareController.GetProgramari)
	router.HandleFunc(utils.FETCH_ALL_APPOINTMENTS_ENDPOINT, programareFetchAllHandler).Methods("GET") // Lists all appointments
	log.Println("[APPOINTMENT] Route GET", utils.FETCH_ALL_APPOINTMENTS_ENDPOINT, "registered.")

	programareFetchByIDHandler := http.HandlerFunc(programareController.GetProgramareByID)
	router.HandleFunc(utils.FETCH_APPOINTMENT_BY_ID_ENDPOINT, programareFetchByIDHandler).Methods("GET") // Get a specific programare by ID
	log.Println("[APPOINTMENT] Route GET", utils.FETCH_APPOINTMENT_BY_ID_ENDPOINT, "registered.")

	programareFetchByDoctorIDHandler := http.HandlerFunc(programareController.GetProgramariByDoctorID)
	router.HandleFunc(utils.FETCH_APPOINTMENTS_BY_DOCTOR_ID_ENDPOINT, programareFetchByDoctorIDHandler).Methods("GET") // Get appointments by doctor ID
	log.Println("[APPOINTMENT] Route GET", utils.FETCH_APPOINTMENTS_BY_DOCTOR_ID_ENDPOINT, "registered.")

	programareFetchByPacientIDHandler := http.HandlerFunc(programareController.GetProgramariByPacientID)
	router.HandleFunc(utils.FETCH_APPOINTMENTS_BY_PACIENT_ID_ENDPOINT, programareFetchByPacientIDHandler).Methods("GET") // Get appointments by pacient ID
	log.Println("[APPOINTMENT] Route GET", utils.FETCH_APPOINTMENTS_BY_PACIENT_ID_ENDPOINT, "registered.")

	programareFetchByDateHandler := http.HandlerFunc(programareController.GetProgramariByDate)
	router.HandleFunc(utils.FETCH_APPOINTMENTS_BY_DATE_ENDPOINT, programareFetchByDateHandler).Methods("GET") // Get appointments by date
	log.Println("[APPOINTMENT] Route GET", utils.FETCH_APPOINTMENTS_BY_DATE_ENDPOINT, "registered.")

	programareFetchByStatusHandler := http.HandlerFunc(programareController.GetProgramariByStatus)
	router.HandleFunc(utils.FETCH_APPOINTMENTS_BY_STATUS_ENDPOINT, programareFetchByStatusHandler).Methods("GET") // Get appointments by status
	log.Println("[APPOINTMENT] Route GET", utils.FETCH_APPOINTMENTS_BY_STATUS_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	programareUpdateByIDHandler := http.HandlerFunc(programareController.UpdateProgramareByID)
	router.Handle(utils.UPDATE_APPOINTMENT_BY_ID_ENDPOINT, middleware.ValidateProgramareInfo(programareUpdateByIDHandler)).Methods("PUT") // Updates a specific programare
	log.Println("[APPOINTMENT] Route PUT", utils.UPDATE_APPOINTMENT_BY_ID_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	programareDeleteByIDHandler := http.HandlerFunc(programareController.DeleteProgramareByID)
	router.Handle(utils.DELETE_APPOINTMENT_BY_ID_ENDPOINT, programareDeleteByIDHandler).Methods("DELETE") // Deletes a programare
	log.Println("[APPOINTMENT] Route DELETE", utils.DELETE_APPOINTMENT_BY_ID_ENDPOINT, "registered.")

	log.Println("[APPOINTMENT] All CRUD routes for Programare entity loaded successfully.")
}
