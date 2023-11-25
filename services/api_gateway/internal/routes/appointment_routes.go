package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/middleware/authorization"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/config"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

func loadAppointmentRoutes(router *mux.Router, gatewayController *controllers.GatewayController, jwtConfig config.JWTConfig) {
	// ---------------------------------------------------------- Create --------------------------------------------------------------
	appointmentCreationHandler := http.HandlerFunc(gatewayController.CreateAppointment)
	router.Handle(utils.CREATE_APPOINTMENT_ENDPOINT, authorization.AdminAndDoctorMiddleware(jwtConfig, appointmentCreationHandler)).Methods("POST")
	log.Println("[APPOINTMENT] Route POST", utils.CREATE_APPOINTMENT_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	appointmentFetchAllHandler := http.HandlerFunc(gatewayController.GetAppointments)
	router.HandleFunc(utils.GET_ALL_APPOINTMENTS_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, appointmentFetchAllHandler)).Methods("GET")
	log.Println("[APPOINTMENT] Route GET", utils.GET_ALL_APPOINTMENTS_ENDPOINT, "registered.")

	appointmentFetchByIDHandler := http.HandlerFunc(gatewayController.GetAppointmentByID)
	router.HandleFunc(utils.GET_APPOINTMENT_BY_ID_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, appointmentFetchByIDHandler)).Methods("GET")
	log.Println("[APPOINTMENT] Route GET", utils.GET_APPOINTMENT_BY_ID_ENDPOINT, "registered.")

	appointmentFetchByDoctorIDHandler := http.HandlerFunc(gatewayController.GetAppointmentsByDoctorID)
	router.HandleFunc(utils.GET_APPOINTMENTS_BY_DOCTOR_ID_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, appointmentFetchByDoctorIDHandler)).Methods("GET")
	log.Println("[APPOINTMENT] Route GET", utils.GET_APPOINTMENTS_BY_DOCTOR_ID_ENDPOINT, "registered.")

	appointmentFetchByPacientIDHandler := http.HandlerFunc(gatewayController.GetAppointmentsByPacientID)
	router.HandleFunc(utils.GET_APPOINTMENTS_BY_PACIENT_ID_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, appointmentFetchByPacientIDHandler)).Methods("GET")
	log.Println("[APPOINTMENT] Route GET", utils.GET_APPOINTMENTS_BY_PACIENT_ID_ENDPOINT, "registered.")

	appointmentFetchByDateHandler := http.HandlerFunc(gatewayController.GetAppointmentsByDate)
	router.HandleFunc(utils.GET_APPOINTMENTS_BY_DATE_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, appointmentFetchByDateHandler)).Methods("GET")
	log.Println("[APPOINTMENT] Route GET", utils.GET_APPOINTMENTS_BY_DATE_ENDPOINT, "registered.")

	appointmentFetchByStatusHandler := http.HandlerFunc(gatewayController.GetAppointmentsByStatus)
	router.HandleFunc(utils.GET_APPOINTMENTS_BY_STATUS_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, appointmentFetchByStatusHandler)).Methods("GET")
	log.Println("[APPOINTMENT] Route GET", utils.GET_APPOINTMENTS_BY_STATUS_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	appointmentUpdateByIDHandler := http.HandlerFunc(gatewayController.UpdateAppointmentByID)
	router.Handle(utils.UPDATE_APPOINTMENT_BY_ID_ENDPOINT, authorization.AdminAndDoctorMiddleware(jwtConfig, appointmentUpdateByIDHandler)).Methods("PUT")
	log.Println("[APPOINTMENT] Route PUT", utils.UPDATE_APPOINTMENT_BY_ID_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	appointmentDeleteByIDHandler := http.HandlerFunc(gatewayController.DeleteAppointmentByID)
	router.Handle(utils.DELETE_APPOINTMENT_BY_ID_ENDPOINT, authorization.AdminAndDoctorMiddleware(jwtConfig, appointmentDeleteByIDHandler)).Methods("DELETE")
	log.Println("[APPOINTMENT] Route DELETE", utils.DELETE_APPOINTMENT_BY_ID_ENDPOINT, "registered.")
}
