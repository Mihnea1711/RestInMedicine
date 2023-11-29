package models

import "github.com/google/uuid"

// ParticipantResponse represents the response from a participant in a transaction
type ParticipantResponse struct {
	ID      uuid.UUID
	Code    int
	Message string
}

// ClientResponse represents the response that is sent to the client after 2pc finishes
type ClientResponse struct {
	Code    int
	Message string
}
