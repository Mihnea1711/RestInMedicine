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
	IDUser int    `json:"idUser"`
}

type ResponseData struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Payload interface{} `json:"payload"`
}

type ActivityData struct {
	IsActive bool `json:"isActive"`
	IDUser   int  `json:"idUser"`
}

type MessageWrapper struct {
	JWT           string `json:"jwt"`
	IDUser        int    `json:"idUser"`
	TransactionID string `json:"transactionID"`
}
