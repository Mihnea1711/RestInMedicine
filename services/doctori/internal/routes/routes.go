package routes

import (
	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/controllers"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	loadCrudRoutes(router)

	return router
}

// loadCrudRoutes loads all the CRUD routes for the Doctor entity
func loadCrudRoutes(r *mux.Router) {
	// Subrouter with the "/doctori" prefix
	doctorRouter := r.PathPrefix("/doctori").Subrouter()

	// Define routes for CRUD operations
	doctorRouter.HandleFunc("/", controllers.GetDoctors).Methods("GET")          // Lists all doctors
	doctorRouter.HandleFunc("/", controllers.CreateDoctor).Methods("POST")       // Creates a new doctor
	doctorRouter.HandleFunc("/{id}", controllers.GetDoctorByID).Methods("GET")   // Fetches a specific doctor by ID
	doctorRouter.HandleFunc("/{id}", controllers.UpdateDoctor).Methods("PUT")    // Updates a specific doctor
	doctorRouter.HandleFunc("/{id}", controllers.DeleteDoctor).Methods("DELETE") // Deletes a doctor
}
