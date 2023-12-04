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

// loadDoctorRoutes loads all the CRUD routes for the Doctor entity
func loadDoctorRoutes(router *mux.Router, gatewayController *controllers.GatewayController, jwtConfig config.JWTConfig) {
	// ---------------------------------------------------------- Create --------------------------------------------------------------
	doctorCreationHandler := http.HandlerFunc(gatewayController.CreateDoctor)
	router.Handle(utils.CREATE_DOCTOR_ENDPOINT, authorization.AdminOnlyMiddleware(jwtConfig, validation.ValidateDoctorData(doctorCreationHandler))).Methods("POST")
	log.Println("[GATEWAY] Route POST", utils.CREATE_DOCTOR_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	doctorFetchAllHandler := http.HandlerFunc(gatewayController.GetDoctors)
	router.HandleFunc(utils.GET_ALL_DOCTORS_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, doctorFetchAllHandler)).Methods("GET")
	log.Println("[GATEWAY] Route GET", utils.GET_ALL_DOCTORS_ENDPOINT, "registered.")

	doctorFetchByEmailHandler := http.HandlerFunc(gatewayController.GetDoctorByEmail)
	router.Handle(utils.GET_DOCTOR_BY_EMAIL_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, doctorFetchByEmailHandler)).Methods("GET")
	log.Println("[GATEWAY] Route GET", utils.GET_DOCTOR_BY_EMAIL_ENDPOINT, "registered.")

	doctorFetchByUserIDHandler := http.HandlerFunc(gatewayController.GetDoctorByUserID)
	router.Handle(utils.GET_DOCTOR_BY_USER_ID_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, doctorFetchByUserIDHandler)).Methods("GET")
	log.Println("[GATEWAY] Route GET", utils.GET_DOCTOR_BY_USER_ID_ENDPOINT, "registered.")

	doctorFetchByIDHandler := http.HandlerFunc(gatewayController.GetDoctorByID)
	router.HandleFunc(utils.GET_DOCTOR_BY_ID_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, doctorFetchByIDHandler)).Methods("GET")
	log.Println("[GATEWAY] Route GET", utils.GET_DOCTOR_BY_ID_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	toggleActivityHandler := http.HandlerFunc(gatewayController.ToggleDoctorActivityByUserID)
	router.Handle(utils.TOGGLE_DOCTOR_ACTIVITY_ENDPOINT, authorization.AdminOnlyMiddleware(jwtConfig, validation.ValidateDoctorActivityData(toggleActivityHandler))).Methods("PATCH")
	log.Println("[PATIENT] Route POST", utils.TOGGLE_DOCTOR_ACTIVITY_ENDPOINT, "registered.")

	doctorUpdateByIDHandler := http.HandlerFunc(gatewayController.UpdateDoctorByID)
	router.Handle(utils.UPDATE_DOCTOR_BY_ID_ENDPOINT, authorization.AdminAndDoctorMiddleware(jwtConfig, validation.ValidateDoctorData(doctorUpdateByIDHandler))).Methods("PUT")
	log.Println("[GATEWAY] Route PUT", utils.UPDATE_DOCTOR_BY_ID_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	doctorDeleteByIDHandler := http.HandlerFunc(gatewayController.DeleteDoctorByID)
	router.Handle(utils.DELETE_DOCTOR_BY_ID_ENDPOINT, authorization.AdminAndDoctorMiddleware(jwtConfig, doctorDeleteByIDHandler)).Methods("DELETE")
	log.Println("[GATEWAY] Route DELETE", utils.DELETE_DOCTOR_BY_ID_ENDPOINT, "registered.")
}
