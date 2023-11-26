package twophasecommit

import (
	"fmt"

	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models"
)

// Pseudo-code for sending prepare and abort messages

func SendPrepareMessage(participants []models.Transactional) []*models.ParticipantResponse {
	var prepareResponses []*models.ParticipantResponse

	for _, participant := range participants {
		response, err := sendMessage(participant, "PREPARE")
		if err != nil {
			fmt.Println("lol")
		}
		prepareResponses = append(prepareResponses, response)
	}

	return prepareResponses
}

func SendAbortMessage(participants []models.Transactional) {
	for _, participant := range participants {
		_, err := sendMessage(participant, "ABORT")
		if err != nil {
			fmt.Println("lol")
		}
	}
}

func SendCommitMessage(participants []models.Transactional) []*models.ParticipantResponse {
	var commitResponses []*models.ParticipantResponse

	for _, participant := range participants {
		response, err := sendMessage(participant, "COMMIT")
		if err != nil {
			fmt.Println("lol")
		}
		commitResponses = append(commitResponses, response)
	}

	return commitResponses
}

func SendRollbackMessage(participants []models.Transactional) {
	for _, participant := range participants {
		_, err := sendMessage(participant, "ROLLBACK")
		if err != nil {
			fmt.Println("lol")
		}

	}
}

// sendMessage is a hypothetical function to send messages to a participant
func sendMessage(participant models.Transactional, messageType string) (*models.ParticipantResponse, error) {
	// Use your messaging infrastructure to send a message to the participant
	// Include the messageType and any other relevant information in the message
	// Return the response received from the participant

	return nil, nil
}
