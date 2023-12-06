package participants

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/utils"
)

// Doctor represents the Doctor module
type Doctor struct {
	models.Participant // Embed Participant struct
	// Additional fields specific to Doctor, if any
}

func NewDoctor(participantID uuid.UUID, participantType models.ParticipantType) *Doctor {
	newDoctor := &Doctor{
		Participant: *models.NewParticipant(participantID, participantType),
	}

	log.Printf("New doctor registered - ID: %s\n", participantID.String())

	return newDoctor
}

// Implement the Transactional interface methods for Participant
func (d *Doctor) Prepare() (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending Doctor prepare request")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), utils.REQUEST_TIMEOUT_MULTIPLIER*time.Second)
	defer cancel() // Make sure to call cancel to release resources associated with the context

	response, status, err := utils.MakeRequest(ctx, http.MethodGet, utils.DOCTOR_HOST, utils.PREPARE_DOCTOR_ENDPOINT, utils.DOCTOR_PORT, nil)
	if err != nil {
		log.Printf("Error making Doctor prepare request: %v", err)
		return nil, err
	}

	log.Printf("[2PC] Doctor prepare request handled successfully. Status Code: %d", status)

	return &models.ParticipantResponse{
		Code:    status,
		Message: response.Message,
	}, nil
}

func (d *Doctor) Commit(userID int) (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending Doctor commit request")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), utils.REQUEST_TIMEOUT_MULTIPLIER*time.Second)
	defer cancel()

	response, status, err := utils.MakeRequest(ctx, http.MethodPatch, utils.DOCTOR_HOST, utils.COMMIT_DOCTOR_ENDPOINT, utils.DOCTOR_PORT, models.ActivityData{
		IsActive: false,
		IDUser:   userID,
	})
	if err != nil {
		log.Printf("Error making Doctor commit request: %v", err)
		return nil, err
	}

	log.Printf("[2PC] Doctor commit request handled successfully. Status Code: %d", status)

	return &models.ParticipantResponse{
		Code:    status,
		Message: response.Message,
	}, nil
}

func (d *Doctor) Abort() (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending Doctor abort request")

	return &models.ParticipantResponse{
		Code:    http.StatusOK,
		Message: "Transaction aborted successfully",
	}, nil
}

func (d *Doctor) Rollback(userID int) (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending Doctor rollback request")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), utils.REQUEST_TIMEOUT_MULTIPLIER*time.Second)
	defer cancel()

	response, status, err := utils.MakeRequest(ctx, http.MethodPatch, utils.DOCTOR_HOST, utils.COMMIT_DOCTOR_ENDPOINT, utils.DOCTOR_PORT, models.ActivityData{
		IsActive: true,
		IDUser:   userID,
	})
	if err != nil {
		log.Printf("Error making Doctor rollback request: %v", err)
		return nil, err
	}

	log.Printf("[2PC] Doctor rollback request handled successfully. Status Code: %d", status)

	return &models.ParticipantResponse{
		Code:    status,
		Message: response.Message,
	}, nil
}
