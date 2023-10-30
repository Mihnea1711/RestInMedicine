package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/database"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/middleware"
)

func SetupRoutes(dbConn database.Database) *mux.Router {
	log.Println("Setting up routes...")
	router := mux.NewRouter()
	router.Use(middleware.RouteLogger)

	doctorController := &controllers.DoctorController{
		DbConn: dbConn,
	}

	loadCrudRoutes(router, doctorController)

	log.Println("Routes setup completed.")
	return router
}

// loadCrudRoutes loads all the CRUD routes for the Doctor entity
func loadCrudRoutes(router *mux.Router, doctorController *controllers.DoctorController) {
	log.Println("Loading CRUD routes for Doctor entity...")

	// ---------------------------------------------------------- Create --------------------------------------------------------------
	doctorCreationHandler := http.HandlerFunc(doctorController.CreateDoctor)
	router.Handle("/doctori", middleware.ValidateDoctorInfo(doctorCreationHandler)).Methods("POST") // Creates a new doctor
	log.Println("Route POST /doctori registered.")

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	doctorFetchAllHandler := http.HandlerFunc(doctorController.GetDoctors)
	router.HandleFunc("/doctori", doctorFetchAllHandler).Methods("GET") // Lists all doctors
	log.Println("Route GET /doctori registered.")

	doctorFetchByIDHandler := http.HandlerFunc(doctorController.GetDoctorByID)
	router.HandleFunc("/doctori/{id}", doctorFetchByIDHandler).Methods("GET") // Lists all doctors
	log.Println("Route GET /doctori/{id} registered.")

	doctorFetchByEmailHandler := http.HandlerFunc(doctorController.GetDoctorByEmail)
	router.Handle("/doctori/email/{email}", middleware.ValidateEmail(doctorFetchByEmailHandler)).Methods("GET")
	log.Println("Route GET /doctori/email/{email} registered.")

	doctorFetchByUserIDHandler := http.HandlerFunc(doctorController.GetDoctorByUserID)
	router.HandleFunc("/doctori/users/{id}", doctorFetchByUserIDHandler).Methods("GET")
	log.Println("Route GET /doctori/users/{id} registered.")

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	doctorUpdateByIDHandler := http.HandlerFunc(doctorController.UpdateDoctorByID)
	router.Handle("/doctori/{id}", middleware.ValidateDoctorInfo(doctorUpdateByIDHandler)).Methods("PUT") // Updates a specific doctor
	log.Println("Route PUT /doctori/{id} registered.")

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	doctorDeleteByIDHandler := http.HandlerFunc(doctorController.DeleteDoctorByID)
	router.Handle("/doctori/{id}", doctorDeleteByIDHandler).Methods("DELETE") // Deletes a doctor
	log.Println("Route DELETE /doctori/{id} registered.")

	log.Println("All CRUD routes for Doctor entity loaded successfully.")
}
