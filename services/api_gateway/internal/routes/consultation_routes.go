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

// loadConsultationRoutes loads all the CRUD routes for the Consultation entity
func loadConsultationRoutes(router *mux.Router, gatewayController *controllers.GatewayController, jwtConfig config.JWTConfig) {
	// ---------------------------------------------------------- Create --------------------------------------------------------------
	consultatieCreationHandler := http.HandlerFunc(gatewayController.CreateConsultation)
	router.Handle(utils.CREATE_CONSULTATION_ENDPOINT, authorization.AdminAndDoctorMiddleware(jwtConfig, validation.ValidateConsultationData(consultatieCreationHandler))).Methods("POST")
	log.Printf("[CONSULTATION] Route POST %s registered.", utils.CREATE_CONSULTATION_ENDPOINT)

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	consultatieFetchAllHandler := http.HandlerFunc(gatewayController.GetConsultations)
	router.HandleFunc(utils.GET_ALL_CONSULTATIONS_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, consultatieFetchAllHandler)).Methods("GET")
	log.Printf("[CONSULTATION] Route GET %s registered.", utils.GET_ALL_CONSULTATIONS_ENDPOINT)

	consultatieFetchByDoctorIDHandler := http.HandlerFunc(gatewayController.GetConsultationsByDoctorID)
	router.HandleFunc(utils.GET_CONSULTATION_BY_DOCTOR_ID_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, consultatieFetchByDoctorIDHandler)).Methods("GET")
	log.Printf("[CONSULTATION] Route GET %s registered.", utils.GET_CONSULTATION_BY_DOCTOR_ID_ENDPOINT)

	consultatieFetchByPacientIDHandler := http.HandlerFunc(gatewayController.GetConsultationsByPatientID)
	router.HandleFunc(utils.GET_CONSULTATION_BY_PATIENT_ID_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, consultatieFetchByPacientIDHandler)).Methods("GET")
	log.Printf("[CONSULTATION] Route GET %s registered.", utils.GET_CONSULTATION_BY_PATIENT_ID_ENDPOINT)

	consultatieFetchByDateHandler := http.HandlerFunc(gatewayController.GetConsultationsByDate)
	router.HandleFunc(utils.GET_CONSULTATION_BY_DATE_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, consultatieFetchByDateHandler)).Methods("GET")
	log.Printf("[CONSULTATION] Route GET %s registered.", utils.GET_CONSULTATION_BY_DATE_ENDPOINT)

	consultatieFetchByIDHandler := http.HandlerFunc(gatewayController.GetConsultationByID)
	router.HandleFunc(utils.GET_CONSULTATION_BY_ID_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, consultatieFetchByIDHandler)).Methods("GET")
	log.Printf("[CONSULTATION] Route GET %s registered.", utils.GET_CONSULTATION_BY_ID_ENDPOINT)

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	consultatieUpdateByIDHandler := http.HandlerFunc(gatewayController.UpdateConsultationByID)
	router.Handle(utils.UPDATE_CONSULTATION_BY_ID_ENDPOINT, authorization.AdminAndDoctorMiddleware(jwtConfig, validation.ValidateConsultationData(consultatieUpdateByIDHandler))).Methods("PUT")
	log.Printf("[CONSULTATION] Route PUT %s registered.", utils.UPDATE_CONSULTATION_BY_ID_ENDPOINT)

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	consultatieDeleteByIDHandler := http.HandlerFunc(gatewayController.DeleteConsultationByID)
	router.Handle(utils.DELETE_CONSULTATION_BY_ID_ENDPOINT, authorization.AdminAndDoctorMiddleware(jwtConfig, consultatieDeleteByIDHandler)).Methods("DELETE")
	log.Printf("[CONSULTATION] Route DELETE %s registered.", utils.DELETE_CONSULTATION_BY_ID_ENDPOINT)
}
