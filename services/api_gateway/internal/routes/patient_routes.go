package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

func loadPatientRoutes(router *mux.Router, gatewayController *controllers.GatewayController) {
	// ---------------------------------------------------------- Create --------------------------------------------------------------
	pacientCreationHandler := http.HandlerFunc(gatewayController.CreatePacient)
	router.Handle(utils.CREATE_PATIENT_ENDPOINT, pacientCreationHandler).Methods("POST")
	log.Printf("[GATEWAY] Route POST %s registered.\n", utils.CREATE_PATIENT_ENDPOINT)

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	pacientFetchAllHandler := http.HandlerFunc(gatewayController.GetPacienti)
	router.HandleFunc(utils.GET_ALL_PATIENTS_ENDPOINT, pacientFetchAllHandler).Methods("GET")
	log.Printf("[GATEWAY] Route GET %s registered.\n", utils.GET_ALL_PATIENTS_ENDPOINT)

	pacientFetchByIDHandler := http.HandlerFunc(gatewayController.GetPacientByID)
	router.HandleFunc(utils.GET_PATIENT_BY_ID_ENDPOINT, pacientFetchByIDHandler).Methods("GET")
	log.Printf("[GATEWAY] Route GET %s registered.\n", utils.GET_PATIENT_BY_ID_ENDPOINT)

	pacientFetchByEmailHandler := http.HandlerFunc(gatewayController.GetPacientByEmail)
	router.Handle(utils.GET_PATIENT_BY_EMAIL_ENDPOINT, pacientFetchByEmailHandler).Methods("GET")
	log.Printf("[GATEWAY] Route GET %s registered.\n", utils.GET_PATIENT_BY_EMAIL_ENDPOINT)

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	pacientUpdateByIDHandler := http.HandlerFunc(gatewayController.UpdatePacientByID)
	router.Handle(utils.UPDATE_PATIENT_BY_ID_ENDPOINT, pacientUpdateByIDHandler).Methods("PUT")
	log.Printf("[GATEWAY] Route PUT %s registered.\n", utils.UPDATE_PATIENT_BY_ID_ENDPOINT)

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	pacientDeleteByIDHandler := http.HandlerFunc(gatewayController.DeletePacientByID)
	router.Handle(utils.DELETE_PATIENT_BY_ID_ENDPOINT, pacientDeleteByIDHandler).Methods("DELETE")
	log.Printf("[GATEWAY] Route DELETE %s registered.\n", utils.DELETE_PATIENT_BY_ID_ENDPOINT)
}
