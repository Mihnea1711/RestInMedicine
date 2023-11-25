package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/middleware/authorization"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/middleware/validation"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/config"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// loadPatientRoutes loads all the CRUD routes for the Patient entity
func loadPatientRoutes(router *mux.Router, gatewayController *controllers.GatewayController, jwtConfig config.JWTConfig) {
	// ---------------------------------------------------------- Create --------------------------------------------------------------
	patientCreationHandler := http.HandlerFunc(gatewayController.CreatePatient)
	router.Handle(utils.CREATE_PATIENT_ENDPOINT, authorization.AdminOnlyMiddleware(jwtConfig, validation.ValidatePatientData(patientCreationHandler))).Methods("POST")
	log.Printf("[GATEWAY] Route POST %s registered.\n", utils.CREATE_PATIENT_ENDPOINT)

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	patientFetchAllHandler := http.HandlerFunc(gatewayController.GetPatients)
	router.HandleFunc(utils.GET_ALL_PATIENTS_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, patientFetchAllHandler)).Methods("GET")
	log.Printf("[GATEWAY] Route GET %s registered.\n", utils.GET_ALL_PATIENTS_ENDPOINT)

	patientFetchByIDHandler := http.HandlerFunc(gatewayController.GetPatientByID)
	router.HandleFunc(utils.GET_PATIENT_BY_ID_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, patientFetchByIDHandler)).Methods("GET")
	log.Printf("[GATEWAY] Route GET %s registered.\n", utils.GET_PATIENT_BY_ID_ENDPOINT)

	patientFetchByEmailHandler := http.HandlerFunc(gatewayController.GetPatientByEmail)
	router.Handle(utils.GET_PATIENT_BY_EMAIL_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, patientFetchByEmailHandler)).Methods("GET")
	log.Printf("[GATEWAY] Route GET %s registered.\n", utils.GET_PATIENT_BY_EMAIL_ENDPOINT)

	patientFetchByUserIDHandler := http.HandlerFunc(gatewayController.GetPatientByUserID)
	router.Handle(utils.GET_PATIENT_BY_USER_ID_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, patientFetchByUserIDHandler)).Methods("GET")
	log.Printf("[GATEWAY] Route GET %s registered.\n", utils.GET_PATIENT_BY_USER_ID_ENDPOINT)

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	patientUpdateByIDHandler := http.HandlerFunc(gatewayController.UpdatePatientByID)
	router.Handle(utils.UPDATE_PATIENT_BY_ID_ENDPOINT, authorization.AdminAndPatientMiddleware(jwtConfig, validation.ValidatePatientData(patientUpdateByIDHandler))).Methods("PUT")
	log.Printf("[GATEWAY] Route PUT %s registered.\n", utils.UPDATE_PATIENT_BY_ID_ENDPOINT)

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	patientDeleteByIDHandler := http.HandlerFunc(gatewayController.DeletePatientByID)
	router.Handle(utils.DELETE_PATIENT_BY_ID_ENDPOINT, authorization.AdminOnlyMiddleware(jwtConfig, patientDeleteByIDHandler)).Methods("DELETE")
	log.Printf("[GATEWAY] Route DELETE %s registered.\n", utils.DELETE_PATIENT_BY_ID_ENDPOINT)
}
