package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

func loadDoctorRoutes(router *mux.Router, gatewayController *controllers.GatewayController) {
	// ---------------------------------------------------------- Create --------------------------------------------------------------
	doctorCreationHandler := http.HandlerFunc(gatewayController.CreateDoctor)
	router.Handle(utils.CREATE_DOCTOR_ENDPOINT, doctorCreationHandler).Methods("POST")
	log.Println("[DOCTOR] Route POST", utils.CREATE_DOCTOR_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	doctorFetchAllHandler := http.HandlerFunc(gatewayController.GetDoctors)
	router.HandleFunc(utils.GET_ALL_DOCTORS_ENDPOINT, doctorFetchAllHandler).Methods("GET")
	log.Println("[DOCTOR] Route GET", utils.GET_ALL_DOCTORS_ENDPOINT, "registered.")

	doctorFetchByIDHandler := http.HandlerFunc(gatewayController.GetDoctorByID)
	router.HandleFunc(utils.GET_DOCTOR_BY_ID_ENDPOINT, doctorFetchByIDHandler).Methods("GET")
	log.Println("[DOCTOR] Route GET", utils.GET_DOCTOR_BY_ID_ENDPOINT, "registered.")

	doctorFetchByEmailHandler := http.HandlerFunc(gatewayController.GetDoctorByEmail)
	router.Handle(utils.GET_DOCTOR_BY_EMAIL_ENDPOINT, doctorFetchByEmailHandler).Methods("GET")
	log.Println("[DOCTOR] Route GET", utils.GET_DOCTOR_BY_EMAIL_ENDPOINT, "registered.")

	doctorFetchByUserIDHandler := http.HandlerFunc(gatewayController.GetDoctorByUserID)
	router.Handle(utils.GET_DOCTOR_BY_USER_ID_ENDPOINT, doctorFetchByUserIDHandler).Methods("GET")
	log.Println("[DOCTOR] Route GET", utils.GET_DOCTOR_BY_USER_ID_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	doctorUpdateByIDHandler := http.HandlerFunc(gatewayController.UpdateDoctorByID)
	router.Handle(utils.UPDATE_DOCTOR_BY_ID_ENDPOINT, doctorUpdateByIDHandler).Methods("PUT")
	log.Println("[DOCTOR] Route PUT", utils.UPDATE_DOCTOR_BY_ID_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	doctorDeleteByIDHandler := http.HandlerFunc(gatewayController.DeleteDoctorByID)
	router.Handle(utils.DELETE_DOCTOR_BY_ID_ENDPOINT, doctorDeleteByIDHandler).Methods("DELETE")
	log.Println("[DOCTOR] Route DELETE", utils.DELETE_DOCTOR_BY_ID_ENDPOINT, "registered.")
}
