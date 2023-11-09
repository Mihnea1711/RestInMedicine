package models

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// MakeAppointmentRequest represents the request payload for making an appointment.
type MakeAppointmentRequest struct {
	IDPatient int    `json:"id_patient"`
	IDDoctor  int    `json:"id_doctor"`
	Date      string `json:"date"`
}

// ProgramConsultationRequest represents the request payload for programming a consultation.
type ProgramConsultationRequest struct {
	IDPatient int    `json:"id_patient"`
	IDDoctor  int    `json:"id_doctor"`
	Date      string `json:"date"`
}
