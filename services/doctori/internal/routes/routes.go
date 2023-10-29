package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/database"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/middleware"
)

func SetupRoutes(dbConn *database.MySQLDatabase) *mux.Router {
	log.Println("Setting up routes...")
	router := mux.NewRouter()
	router.Use(middleware.Logger)

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
	/*
		// Subrouter with the "/doctori" prefix
		crudRouter := r.PathPrefix("/doctori").Subrouter()
	*/

	// Define routes for CRUD operations
	// Define routes for CRUD operations
	doctorCreationHandler := http.HandlerFunc(doctorController.CreateDoctor)
	router.Handle("/doctori", middleware.ValidateDoctorCreation(doctorCreationHandler)).Methods("POST") // Creates a new doctor
	log.Println("Route POST /doctori registered.")

	doctorFetchAllHandler := http.HandlerFunc(doctorController.GetDoctors)
	router.HandleFunc("/doctori", doctorFetchAllHandler).Methods("GET") // Lists all doctors
	log.Println("Route GET /doctori registered.")

	router.HandleFunc("/doctori/{id}", controllers.GetDoctorByID).Methods("GET") // Fetches a specific doctor by ID
	log.Println("Route GET /doctori/{id} registered.")

	router.HandleFunc("/doctori/{id}", controllers.UpdateDoctor).Methods("PUT") // Updates a specific doctor
	log.Println("Route PUT /doctori/{id} registered.")

	router.HandleFunc("/doctori/{id}", controllers.DeleteDoctor).Methods("DELETE") // Deletes a doctor
	log.Println("Route DELETE /doctori/{id} registered.")

	log.Println("All CRUD routes for Doctor entity loaded successfully.")
}
