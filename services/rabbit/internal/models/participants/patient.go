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

// Patient represents the Patient module
type Patient struct {
	models.Participant // Embed Participant struct
	// Additional fields specific to Patient, if any
}

func NewPatient(participantID uuid.UUID, participantType models.ParticipantType) *Patient {
	return &Patient{
		Participant: *models.NewParticipant(participantID, participantType),
	}
}

// Implement the Transactional interface methods for Participant
func (p *Patient) Prepare() (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending patient prepare request")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), utils.REQUEST_TIMEOUT_MULTIPLIER*time.Second)
	defer cancel() // Make sure to call cancel to release resources associated with the context

	response, status, err := utils.MakeRequest(ctx, http.MethodGet, utils.PATIENT_HOST, utils.PREPARE_PATIENT_ENDPOINT, utils.PATIENT_PORT, nil)
	if err != nil {
		log.Printf("Error making patient prepare request: %v", err)
		return nil, err
	}

	return &models.ParticipantResponse{
		Code:    status,
		Message: response.Message,
	}, nil
}

func (p *Patient) Commit(userID int) (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending patient commit request")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), utils.REQUEST_TIMEOUT_MULTIPLIER*time.Second)
	defer cancel()

	response, status, err := utils.MakeRequest(ctx, http.MethodPatch, utils.PATIENT_HOST, utils.COMMIT_PATIENT_ENDPOINT, utils.PATIENT_PORT, models.ActivityData{
		IsActive: false,
		IDUser:   userID,
	})
	if err != nil {
		log.Printf("Error making patient commit request: %v", err)
		return nil, err
	}

	return &models.ParticipantResponse{
		Code:    status,
		Message: response.Message,
	}, nil
}

func (p *Patient) Abort() (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending patient abort request")

	return &models.ParticipantResponse{
		Code:    http.StatusOK,
		Message: "Transaction aborted successfully",
	}, nil
}

func (p *Patient) Rollback() (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending patient rollback request")

	return &models.ParticipantResponse{
		Code:    http.StatusOK,
		Message: "Transaction rolled back successfully",
	}, nil
}

func (p *Patient) Compensate() error {
	return nil
}
