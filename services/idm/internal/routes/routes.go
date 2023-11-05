package routes

import (
	"log"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/idm/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/idm/internal/database"
	"github.com/mihnea1711/POS_Project/services/idm/internal/database/redis"
	"github.com/mihnea1711/POS_Project/services/idm/internal/middleware"
)

func SetupRoutes(dbConn database.Database, rdb *redis.RedisClient) *mux.Router {
	log.Println("[PROGRAMARE] Setting up rate limiter...")
	rateLimiter := middleware.NewRedisRateLimiter(rdb, 10, time.Minute) // Here, I'm allowing 10 requests per minute.

	log.Println("[PROGRAMARE] Setting up routes...")
	router := mux.NewRouter()
	router.Use(rateLimiter.Limit)
	router.Use(middleware.RouteLogger)

	programariController := &controllers.ProgramareController{
		DbConn: dbConn,
	}

	loadCrudRoutes(router, programariController)

	log.Println("[PROGRAMARE] Routes setup completed.")
	return router
}

// loadCrudRoutes loads all the CRUD routes for the Doctor entity
func loadCrudRoutes(router *mux.Router, programareController *controllers.ProgramareController) {
	log.Println("[PROGRAMARE] Loading CRUD routes for Programare entity...")

	// // ---------------------------------------------------------- Create --------------------------------------------------------------
	// programareCreationHandler := http.HandlerFunc(programareController.CreateProgramare)
	// router.Handle("/programari", middleware.ValidateProgramareInfo(programareCreationHandler)).Methods("POST") // Creates a new programare
	// log.Println("[PROGRAMARE] Route POST /programari registered.")

	// // ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	// // implement pagination for these
	// programareFetchAllHandler := http.HandlerFunc(programareController.GetProgramari)
	// router.HandleFunc("/programari", programareFetchAllHandler).Methods("GET") // Lists all programari
	// log.Println("[PROGRAMARE] Route GET /programari registered.")

	// programareFetchByIDHandler := http.HandlerFunc(programareController.GetProgramareByID)
	// router.HandleFunc("/programari/{id}", programareFetchByIDHandler).Methods("GET") // Get a specific programare by ID
	// log.Println("[PROGRAMARE] Route GET /programari/{id} registered.")

	// programareFetchByDoctorIDHandler := http.HandlerFunc(programareController.GetProgramariByDoctorID)
	// router.HandleFunc("/programari/doctor/{id}", programareFetchByDoctorIDHandler).Methods("GET") // Get programari by doctor ID
	// log.Println("[PROGRAMARE] Route GET /programari/doctor/{id} registered.")

	// programareFetchByPacientIDHandler := http.HandlerFunc(programareController.GetProgramariByPacientID)
	// router.HandleFunc("/programari/pacient/{id}", programareFetchByPacientIDHandler).Methods("GET") // Get programari by pacient ID
	// log.Println("[PROGRAMARE] Route GET /programari/pacient/{id} registered.")

	// programareFetchByDateHandler := http.HandlerFunc(programareController.GetProgramariByDate)
	// router.HandleFunc("/programari/date/{date}", programareFetchByDateHandler).Methods("GET") // Get programari by date
	// log.Println("[PROGRAMARE] Route GET /programari/date/{date} registered.")

	// programareFetchByStatusHandler := http.HandlerFunc(programareController.GetProgramariByStatus)
	// router.HandleFunc("/programari/status/{status}", programareFetchByStatusHandler).Methods("GET") // Get programari by status
	// log.Println("[PROGRAMARE] Route GET /programari/status/{status} registered.")

	// // ---------------------------------------------------------- Update --------------------------------------------------------------
	// programareUpdateByIDHandler := http.HandlerFunc(programareController.UpdateProgramareByID)
	// router.Handle("/programari/{id}", middleware.ValidateProgramareInfo(programareUpdateByIDHandler)).Methods("PUT") // Updates a specific programare
	// log.Println("[PROGRAMARE] Route PUT /programari/{id} registered.")

	// // ---------------------------------------------------------- Delete --------------------------------------------------------------
	// programareDeleteByIDHandler := http.HandlerFunc(programareController.DeleteProgramareByID)
	// router.Handle("/programari/{id}", programareDeleteByIDHandler).Methods("DELETE") // Deletes a programare
	// log.Println("[PROGRAMARE] Route DELETE /programari/{id} registered.")

	log.Println("[PROGRAMARE] All CRUD routes for Programare entity loaded successfully.")
}
