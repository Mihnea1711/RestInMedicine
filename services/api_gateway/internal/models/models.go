package models

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
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

type UserData struct {
	Username string `json:"username"`
	// other if needed
}

type PasswordData struct {
	Password string `json:"password"`
}

type RoleData struct {
	Role string `json:"role"`
}
