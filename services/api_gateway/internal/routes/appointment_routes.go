package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/controllers"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

func loadAppointmentRoutes(router *mux.Router, gatewayController *controllers.GatewayController) {
	// ---------------------------------------------------------- Create --------------------------------------------------------------
	appointmentCreationHandler := http.HandlerFunc(gatewayController.CreateAppointment)
	router.Handle(utils.CREATE_APPOINTMENT_ENDPOINT, appointmentCreationHandler).Methods("POST") // Creates a new appointment
	log.Println("[APPOINTMENT] Route POST", utils.CREATE_APPOINTMENT_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Retrieve --------------------------------------------------------------
	appointmentFetchAllHandler := http.HandlerFunc(gatewayController.GetAppointments)
	router.HandleFunc(utils.GET_ALL_APPOINTMENTS_ENDPOINT, appointmentFetchAllHandler).Methods("GET") // Lists all appointments
	log.Println("[APPOINTMENT] Route GET", utils.GET_ALL_APPOINTMENTS_ENDPOINT, "registered.")

	appointmentFetchByIDHandler := http.HandlerFunc(gatewayController.GetAppointmentByID)
	router.HandleFunc(utils.GET_APPOINTMENT_BY_ID_ENDPOINT, appointmentFetchByIDHandler).Methods("GET") // Get a specific appointment by ID
	log.Println("[APPOINTMENT] Route GET", utils.GET_APPOINTMENT_BY_ID_ENDPOINT, "registered.")

	appointmentFetchByDoctorIDHandler := http.HandlerFunc(gatewayController.GetAppointmentsByDoctorID)
	router.HandleFunc(utils.GET_APPOINTMENTS_BY_DOCTOR_ID_ENDPOINT, appointmentFetchByDoctorIDHandler).Methods("GET") // Get appointments by doctor ID
	log.Println("[APPOINTMENT] Route GET", utils.GET_APPOINTMENTS_BY_DOCTOR_ID_ENDPOINT, "registered.")

	appointmentFetchByPacientIDHandler := http.HandlerFunc(gatewayController.GetAppointmentsByPacientID)
	router.HandleFunc(utils.GET_APPOINTMENTS_BY_PACIENT_ID_ENDPOINT, appointmentFetchByPacientIDHandler).Methods("GET") // Get appointments by pacient ID
	log.Println("[APPOINTMENT] Route GET", utils.GET_APPOINTMENTS_BY_PACIENT_ID_ENDPOINT, "registered.")

	appointmentFetchByDateHandler := http.HandlerFunc(gatewayController.GetAppointmentsByDate)
	router.HandleFunc(utils.GET_APPOINTMENTS_BY_DATE_ENDPOINT, appointmentFetchByDateHandler).Methods("GET") // Get appointments by date
	log.Println("[APPOINTMENT] Route GET", utils.GET_APPOINTMENTS_BY_DATE_ENDPOINT, "registered.")

	appointmentFetchByStatusHandler := http.HandlerFunc(gatewayController.GetAppointmentsByStatus)
	router.HandleFunc(utils.GET_APPOINTMENTS_BY_STATUS_ENDPOINT, appointmentFetchByStatusHandler).Methods("GET") // Get appointments by status
	log.Println("[APPOINTMENT] Route GET", utils.GET_APPOINTMENTS_BY_STATUS_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Update --------------------------------------------------------------
	appointmentUpdateByIDHandler := http.HandlerFunc(gatewayController.UpdateAppointmentByID)
	router.Handle(utils.UPDATE_APPOINTMENT_BY_ID_ENDPOINT, appointmentUpdateByIDHandler).Methods("PUT") // Updates a specific appointment
	log.Println("[APPOINTMENT] Route PUT", utils.UPDATE_APPOINTMENT_BY_ID_ENDPOINT, "registered.")

	// ---------------------------------------------------------- Delete --------------------------------------------------------------
	appointmentDeleteByIDHandler := http.HandlerFunc(gatewayController.DeleteAppointmentByID)
	router.Handle(utils.DELETE_APPOINTMENT_BY_ID_ENDPOINT, appointmentDeleteByIDHandler).Methods("DELETE") // Deletes a appointment
	log.Println("[APPOINTMENT] Route DELETE", utils.DELETE_APPOINTMENT_BY_ID_ENDPOINT, "registered.")

}
