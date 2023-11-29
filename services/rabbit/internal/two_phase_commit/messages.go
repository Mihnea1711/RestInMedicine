package twophasecommit

import (
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models"
)

func SendPrepareMessage(participants []models.Participant) ([]*models.ParticipantResponse, error) {
	var prepareResponses []*models.ParticipantResponse

	for _, participant := range participants {
		response, err := sendMessage(participant, "PREPARE")
		if err != nil {
			// Handle the error, e.g., log it
			log.Printf("Error sending PREPARE message: %v\n", err)
			continue
		}
		prepareResponses = append(prepareResponses, response)
	}

	return prepareResponses, nil
}

func SendAbortMessage(participants []models.Participant) ([]*models.ParticipantResponse, error) {
	var abortResponses []*models.ParticipantResponse

	for _, participant := range participants {
		response, err := sendMessage(participant, "ABORT")
		if err != nil {
			// Handle the error, e.g., log it
			log.Printf("Error sending ABORT message: %v\n", err)
			continue
		}
		abortResponses = append(abortResponses, response)
	}

	return abortResponses, nil
}

func SendCommitMessage(participants []models.Participant) ([]*models.ParticipantResponse, error) {
	var commitResponses []*models.ParticipantResponse

	for _, participant := range participants {
		response, err := sendMessage(participant, "COMMIT")
		if err != nil {
			// Handle the error, e.g., log it
			log.Printf("Error sending COMMIT message: %v\n", err)
			continue
		}
		commitResponses = append(commitResponses, response)
	}

	return commitResponses, nil
}

func SendRollbackMessage(participants []models.Participant) ([]*models.ParticipantResponse, error) {
	var rollbackResponses []*models.ParticipantResponse

	for _, participant := range participants {
		response, err := sendMessage(participant, "ROLLBACK")
		if err != nil {
			// Handle the error, e.g., log it
			log.Printf("Error sending ROLLBACK message: %v\n", err)
			continue
		}
		rollbackResponses = append(rollbackResponses, response)
	}

	return rollbackResponses, nil
}

// sendMessage is a hypothetical function to send messages to a participant
func sendMessage(participant models.Participant, messageType string) (*models.ParticipantResponse, error) {
	var response *models.ParticipantResponse
	var err error

	switch messageType {
	case "PREPARE":
		response, err = participant.Prepare()
	case "COMMIT":
		response, err = participant.Commit()
	case "ROLLBACK", "ABORT":
		response, err = participant.Rollback()
	default:
		err = fmt.Errorf("unsupported message type: %s", messageType)
	}

	// Construct the response based on the outcome of the transactional methods
	if err != nil {
		log.Println("An error occured while trying to send the message")
		return nil, err
	}

	return response, nil
}
