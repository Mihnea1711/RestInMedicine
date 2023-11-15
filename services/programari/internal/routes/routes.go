package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/programari/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/programari/internal/database"
	"github.com/mihnea1711/POS_Project/services/programari/internal/database/redis"
	"github.com/mihnea1711/POS_Project/services/programari/internal/middleware"
)

func SetupRoutes(dbConn database.Database, rdb *redis.RedisClient) *mux.Router {
	log.Println("[APPOINTMENT] Setting up rate limiter...")
	rateLimiter := middleware.NewRedisRateLimiter(rdb.GetClient(), 10, time.Minute) // Here, I'm allowing 10 requests per minute.

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
	log.Println("[APPOINTMENT] Loading CRUD routes for Programare entity...")

	// ---------------------------------------------------------- Create --------------------------------------------------------------
	programareCreationHandler := http.HandlerFunc(programareController.CreateProgramare)
	router.Handle("/appointments", middleware.ValidateProgramareInfo(programareCreationHandler)).Methods("POST") // Creates a new programare
	log.Println("[APPOINTMENT] Route POST /appointments registered.")

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	// implement pagination for these
	programareFetchAllHandler := http.HandlerFunc(programareController.GetProgramari)
	router.HandleFunc("/appointments", programareFetchAllHandler).Methods("GET") // Lists all appointments
	log.Println("[APPOINTMENT] Route GET /appointments registered.")

	programareFetchByIDHandler := http.HandlerFunc(programareController.GetProgramareByID)
	router.HandleFunc("/appointments/{id}", programareFetchByIDHandler).Methods("GET") // Get a specific programare by ID
	log.Println("[APPOINTMENT] Route GET /appointments/{id} registered.")

	programareFetchByDoctorIDHandler := http.HandlerFunc(programareController.GetProgramariByDoctorID)
	router.HandleFunc("/appointments/doctor/{id}", programareFetchByDoctorIDHandler).Methods("GET") // Get appointments by doctor ID
	log.Println("[APPOINTMENT] Route GET /appointments/doctor/{id} registered.")

	programareFetchByPacientIDHandler := http.HandlerFunc(programareController.GetProgramariByPacientID)
	router.HandleFunc("/appointments/pacient/{id}", programareFetchByPacientIDHandler).Methods("GET") // Get appointments by pacient ID
	log.Println("[APPOINTMENT] Route GET /appointments/pacient/{id} registered.")

	programareFetchByDateHandler := http.HandlerFunc(programareController.GetProgramariByDate)
	router.HandleFunc("/appointments/date/{date}", programareFetchByDateHandler).Methods("GET") // Get appointments by date
	log.Println("[APPOINTMENT] Route GET /appointments/date/{date} registered.")

	programareFetchByStatusHandler := http.HandlerFunc(programareController.GetProgramariByStatus)
	router.HandleFunc("/appointments/status/{status}", programareFetchByStatusHandler).Methods("GET") // Get appointments by status
	log.Println("[APPOINTMENT] Route GET /appointments/status/{status} registered.")

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	programareUpdateByIDHandler := http.HandlerFunc(programareController.UpdateProgramareByID)
	router.Handle("/appointments/{id}", middleware.ValidateProgramareInfo(programareUpdateByIDHandler)).Methods("PUT") // Updates a specific programare
	log.Println("[APPOINTMENT] Route PUT /appointments/{id} registered.")

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	programareDeleteByIDHandler := http.HandlerFunc(programareController.DeleteProgramareByID)
	router.Handle("/appointments/{id}", programareDeleteByIDHandler).Methods("DELETE") // Deletes a programare
	log.Println("[APPOINTMENT] Route DELETE /appointments/{id} registered.")

	log.Println("[APPOINTMENT] All CRUD routes for Programare entity loaded successfully.")
}
