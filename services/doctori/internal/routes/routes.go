package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/middleware"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.Logger)

	loadCrudRoutes(router)

	return router
}

// loadCrudRoutes loads all the CRUD routes for the Doctor entity
func loadCrudRoutes(router *mux.Router) {
	// Subrouter with the "/doctori" prefix
	// crudRouter := r.PathPrefix("/doctori").Subrouter()

	// Define routes for CRUD operations
	doctorCreationHandler := http.HandlerFunc(controllers.CreateDoctor)
	router.Handle("/doctori", middleware.ValidateDoctorCreation(doctorCreationHandler)).Methods("POST") // Creates a new doctor

	router.HandleFunc("/doctori", controllers.GetDoctors).Methods("GET")           // Lists all doctors
	router.HandleFunc("/doctori/{id}", controllers.GetDoctorByID).Methods("GET")   // Fetches a specific doctor by ID
	router.HandleFunc("/doctori/{id}", controllers.UpdateDoctor).Methods("PUT")    // Updates a specific doctor
	router.HandleFunc("/doctori/{id}", controllers.DeleteDoctor).Methods("DELETE") // Deletes a doctor
}
