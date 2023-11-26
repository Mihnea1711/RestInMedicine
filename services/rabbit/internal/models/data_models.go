package models

// ParticipantResponse represents the response from a participant in a transaction
type ParticipantResponse struct {
	ParticipantID string
	Success       bool
	// Include other relevant information in the response if needed
}

// Participant represents a module that participates in a distributed transaction
type ParticipantData struct {
	// Include other relevant fields if needed
	ID   string
	Name string
}
