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

// loadAppointmentRoutes loads all the CRUD routes for the Appointment entity
func loadAppointmentRoutes(router *mux.Router, gatewayController *controllers.GatewayController, jwtConfig config.JWTConfig) {
	// ---------------------------------------------------------- Create --------------------------------------------------------------
	appointmentCreationHandler := http.HandlerFunc(gatewayController.CreateAppointment)
	router.Handle(utils.CREATE_APPOINTMENT_ENDPOINT, authorization.AdminAndDoctorMiddleware(jwtConfig, validation.ValidateAppointmentData(appointmentCreationHandler))).Methods("POST")
	log.Println("[GATEWAY] Route POST", utils.CREATE_APPOINTMENT_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	appointmentFetchAllHandler := http.HandlerFunc(gatewayController.GetAppointments)
	router.HandleFunc(utils.GET_ALL_APPOINTMENTS_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, appointmentFetchAllHandler)).Methods("GET")
	log.Println("[GATEWAY] Route GET", utils.GET_ALL_APPOINTMENTS_ENDPOINT, "registered.")

	appointmentFetchByIDHandler := http.HandlerFunc(gatewayController.GetAppointmentByID)
	router.HandleFunc(utils.GET_APPOINTMENT_BY_ID_ENDPOINT, authorization.AllRolesMiddleware(jwtConfig, appointmentFetchByIDHandler)).Methods("GET")
	log.Println("[GATEWAY] Route GET", utils.GET_APPOINTMENT_BY_ID_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	appointmentUpdateByIDHandler := http.HandlerFunc(gatewayController.UpdateAppointmentByID)
	router.Handle(utils.UPDATE_APPOINTMENT_BY_ID_ENDPOINT, authorization.AdminAndDoctorMiddleware(jwtConfig, validation.ValidateAppointmentData(appointmentUpdateByIDHandler))).Methods("PUT")
	log.Println("[GATEWAY] Route PUT", utils.UPDATE_APPOINTMENT_BY_ID_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	appointmentDeleteByIDHandler := http.HandlerFunc(gatewayController.DeleteAppointmentByID)
	router.Handle(utils.DELETE_APPOINTMENT_BY_ID_ENDPOINT, authorization.AdminAndDoctorMiddleware(jwtConfig, appointmentDeleteByIDHandler)).Methods("DELETE")
	log.Println("[GATEWAY] Route DELETE", utils.DELETE_APPOINTMENT_BY_ID_ENDPOINT, "registered.")
}
