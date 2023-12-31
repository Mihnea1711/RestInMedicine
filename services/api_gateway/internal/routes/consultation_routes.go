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
	log.Printf("[GATEWAY] Route POST %s registered.", utils.CREATE_CONSULTATION_ENDPOINT)

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	consultatieFetchAllHandler := http.HandlerFunc(gatewayController.GetConsultations)
	router.HandleFunc(utils.GET_ALL_CONSULTATIONS_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, consultatieFetchAllHandler)).Methods("GET")
	log.Printf("[GATEWAY] Route GET %s registered.", utils.GET_ALL_CONSULTATIONS_ENDPOINT)

	consultatieFetchByIDHandler := http.HandlerFunc(gatewayController.GetConsultationByID)
	router.HandleFunc(utils.GET_CONSULTATION_BY_ID_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, consultatieFetchByIDHandler)).Methods("GET")
	log.Printf("[GATEWAY] Route GET %s registered.", utils.GET_CONSULTATION_BY_ID_ENDPOINT)

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	consultatieUpdateByIDHandler := http.HandlerFunc(gatewayController.UpdateConsultationByID)
	router.Handle(utils.UPDATE_CONSULTATION_BY_ID_ENDPOINT, authorization.AdminAndDoctorMiddleware(jwtConfig, validation.ValidateConsultationData(consultatieUpdateByIDHandler))).Methods("PUT")
	log.Printf("[GATEWAY] Route PUT %s registered.", utils.UPDATE_CONSULTATION_BY_ID_ENDPOINT)

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	consultatieDeleteByIDHandler := http.HandlerFunc(gatewayController.DeleteConsultationByID)
	router.Handle(utils.DELETE_CONSULTATION_BY_ID_ENDPOINT, authorization.AdminAndDoctorMiddleware(jwtConfig, consultatieDeleteByIDHandler)).Methods("DELETE")
	log.Printf("[GATEWAY] Route DELETE %s registered.", utils.DELETE_CONSULTATION_BY_ID_ENDPOINT)
}
