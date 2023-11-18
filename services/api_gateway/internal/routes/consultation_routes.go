package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

func loadConsultationRoutes(router *mux.Router, gatewayController *controllers.GatewayController) {
	// ---------------------------------------------------------- Create --------------------------------------------------------------
	consultatieCreationHandler := http.HandlerFunc(gatewayController.CreateConsultation)
	router.Handle(utils.CREATE_CONSULTATION_ENDPOINT, consultatieCreationHandler).Methods("POST")
	log.Printf("[CONSULTATION] Route POST %s registered.", utils.CREATE_CONSULTATION_ENDPOINT)

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	consultatieFetchAllHandler := http.HandlerFunc(gatewayController.GetConsultations)
	router.HandleFunc(utils.GET_ALL_CONSULTATIONS_ENDPOINT, consultatieFetchAllHandler).Methods("GET")
	log.Printf("[CONSULTATION] Route GET %s registered.", utils.GET_ALL_CONSULTATIONS_ENDPOINT)

	consultatieFetchByDoctorIDHandler := http.HandlerFunc(gatewayController.GetConsultationsByDoctorID)
	router.HandleFunc(utils.GET_CONSULTATION_BY_DOCTOR_ID_ENDPOINT, consultatieFetchByDoctorIDHandler).Methods("GET")
	log.Printf("[CONSULTATION] Route GET %s registered.", utils.GET_CONSULTATION_BY_DOCTOR_ID_ENDPOINT)

	consultatieFetchByPacientIDHandler := http.HandlerFunc(gatewayController.GetConsultationsByPacientID)
	router.HandleFunc(utils.GET_CONSULTATION_BY_PACIENT_ID_ENDPOINT, consultatieFetchByPacientIDHandler).Methods("GET")
	log.Printf("[CONSULTATION] Route GET %s registered.", utils.GET_CONSULTATION_BY_PACIENT_ID_ENDPOINT)

	consultatieFetchByDateHandler := http.HandlerFunc(gatewayController.GetConsultationsByDate)
	router.HandleFunc(utils.GET_CONSULTATION_BY_DATE_ENDPOINT, consultatieFetchByDateHandler).Methods("GET")
	log.Printf("[CONSULTATION] Route GET %s registered.", utils.GET_CONSULTATION_BY_DATE_ENDPOINT)

	// filteredConsultationHandler := http.HandlerFunc(gatewayController.GetFilteredConsultations)
	// router.HandleFunc(utils.FILTER_CONSULTATIONS_ENDPOINT, filteredConsultationHandler).Methods("GET")
	// log.Printf("[CONSULTATION] Route GET %s registered.", utils.FILTER_CONSULTATIONS_ENDPOINT)

	consultatieFetchByIDHandler := http.HandlerFunc(gatewayController.GetConsultationByID)
	router.HandleFunc(utils.GET_CONSULTATION_BY_ID_ENDPOINT, consultatieFetchByIDHandler).Methods("GET")
	log.Printf("[CONSULTATION] Route GET %s registered.", utils.GET_CONSULTATION_BY_ID_ENDPOINT)

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	consultatieUpdateByIDHandler := http.HandlerFunc(gatewayController.UpdateConsultationByID)
	router.Handle(utils.UPDATE_CONSULTATION_BY_ID_ENDPOINT, consultatieUpdateByIDHandler).Methods("PUT")
	log.Printf("[CONSULTATION] Route PUT %s registered.", utils.UPDATE_CONSULTATION_BY_ID_ENDPOINT)

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	consultatieDeleteByIDHandler := http.HandlerFunc(gatewayController.DeleteConsultationByID)
	router.Handle(utils.DELETE_CONSULTATION_BY_ID_ENDPOINT, consultatieDeleteByIDHandler).Methods("DELETE")
	log.Printf("[CONSULTATION] Route DELETE %s registered.", utils.DELETE_CONSULTATION_BY_ID_ENDPOINT)

}
