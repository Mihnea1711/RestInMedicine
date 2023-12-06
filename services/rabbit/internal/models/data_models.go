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

type DeleteMessageData struct {
	JWT    string `json:"jwt"`
	IDUser int    `json:"id_user"`
}

type ResponseData struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Payload interface{} `json:"payload"`
}

type ActivityData struct {
	IsActive bool `json:"is_active"`
	IDUser   int  `json:"id_user"`
}

type MessageWrapper struct {
	JWT           string `json:"jwt"`
	IDUser        int    `json:"id_user"`
	TransactionID string `json:"transaction_id"`
}
