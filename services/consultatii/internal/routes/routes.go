package routes

import (
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

func SetupRoutes(dbConn database.Database, rdb *redis.RedisClient) *mux.Router {
	log.Println("[CONSULTATIE] Setting up rate limiter...")
	rateLimiter := middleware.NewRedisRateLimiter(rdb, utils.REQUEST_RATE, utils.REQUEST_WINDOW_DURATION_MULTIPLIER*time.Minute) // Here, I'm allowing 10 requests per minute * MULTIPLIER.

	log.Println("[CONSULTATIE] Setting up routes...")
	router := mux.NewRouter()
	router.Use(rateLimiter.Limit)
	router.Use(middleware.RouteLogger)

	consultatieController := &controllers.ConsultatieController{
		DbConn: dbConn,
	}

	loadCrudRoutes(router, consultatieController)

	log.Println("[CONSULTATIE] Routes setup completed.")
	return router
}

// loadCrudRoutes loads all the CRUD routes for the consultatie entity
func loadCrudRoutes(router *mux.Router, consultatieController *controllers.ConsultatieController) {
	log.Println("[CONSULTATIE] Loading CRUD routes for consultatie entity...")

	// // ---------------------------------------------------------- Create --------------------------------------------------------------
	consultatieCreationHandler := http.HandlerFunc(consultatieController.CreateConsultatie)
	router.Handle(utils.INSERT_CONSULTATIE_ENDPOINT, middleware.ValidateConsultatieInfo(consultatieCreationHandler)).Methods("POST") // Creates a new consultatie
	log.Print("[CONSULTATIE] Route POST /consultatii registered.")

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	consultatieFetchAllHandler := http.HandlerFunc(consultatieController.GetConsultatii)
	router.HandleFunc(utils.FETCH_ALL_CONSULTATII_ENDPOINT, consultatieFetchAllHandler).Methods("GET") // Lists all consultaties
	log.Printf("[CONSULTATIE] Route GET %s registered.", utils.FETCH_ALL_CONSULTATII_ENDPOINT)

	consultatieFetchByDoctorIDHandler := http.HandlerFunc(consultatieController.GetConsultatiiByDoctorID)
	router.HandleFunc(utils.FETCH_CONSULTATIE_BY_DOCTOR_ID_ENDPOINT, consultatieFetchByDoctorIDHandler).Methods("GET") // Get consultatii by doctor ID
	log.Printf("[CONSULTATIE] Route GET %s registered.", utils.FETCH_CONSULTATIE_BY_DOCTOR_ID_ENDPOINT)

	consultatieFetchByPacientIDHandler := http.HandlerFunc(consultatieController.GetConsultatiiByPacientID)
	router.HandleFunc(utils.FETCH_CONSULTATIE_BY_PACIENT_ID_ENDPOINT, consultatieFetchByPacientIDHandler).Methods("GET") // Get consultatii by pacient ID
	log.Printf("[CONSULTATIE] Route GET %s registered.", utils.FETCH_CONSULTATIE_BY_PACIENT_ID_ENDPOINT)

	consultatieFetchByDateHandler := http.HandlerFunc(consultatieController.GetConsultatiiByDate)
	router.HandleFunc(utils.FETCH_CONSULTATIE_BY_DATE_ENDPOINT, consultatieFetchByDateHandler).Methods("GET") // Get consultatii by date
	log.Printf("[CONSULTATIE] Route GET %s registered.", utils.FETCH_CONSULTATIE_BY_DATE_ENDPOINT)

	consultatieFetchByIDHandler := http.HandlerFunc(consultatieController.GetConsultatieByID)
	router.HandleFunc(utils.FETCH_CONSULTATIE_BY_ID_ENDPOINT, consultatieFetchByIDHandler).Methods("GET") // Get consultatii by consultatie ID
	log.Printf("[CONSULTATIE] Route GET %s registered.", utils.FETCH_CONSULTATIE_BY_ID_ENDPOINT)

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	consultatieUpdateByIDHandler := http.HandlerFunc(consultatieController.UpdateConsultatieByID)
	router.Handle(utils.UPDATE_CONSULTATIE_BY_ID_ENDPOINT, middleware.ValidateConsultatieInfo(consultatieUpdateByIDHandler)).Methods("PUT") // Updates a specific consultatie
	log.Printf("[CONSULTATIE] Route PUT %s registered.", utils.UPDATE_CONSULTATIE_BY_ID_ENDPOINT)

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	consultatieDeleteByIDHandler := http.HandlerFunc(consultatieController.DeleteConsultatieByID)
	router.Handle(utils.DELETE_CONSULTATIE_BY_ID_ENDPOINT, consultatieDeleteByIDHandler).Methods("DELETE") // Deletes a consultatie
	log.Printf("[CONSULTATIE] Route DELETE %s registered.", utils.DELETE_CONSULTATIE_BY_ID_ENDPOINT)

	log.Println("[CONSULTATIE] All CRUD routes for consultatie entity loaded successfully.")
}
