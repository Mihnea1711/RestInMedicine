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
	// Implement preparation logic

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), utils.REQUEST_TIMEOUT_MULTIPLIER*time.Second)
	defer cancel() // Make sure to call cancel to release resources associated with the context

	response, status, err := utils.MakeRequest(ctx, http.MethodGet, utils.PATIENT_HOST, utils.PATIENT_PORT, utils.PREPARE_PATIENT_ENDPOINT)
	if err != nil {
		// smth
		log.Printf("Error making req: %v", err)
	}

	return &models.ParticipantResponse{
		ID:      response.ID,
		Code:    status,
		Message: response.Message,
	}, nil
}

func (p *Patient) Commit() (*models.ParticipantResponse, error) {
	// Implement commit logic
	return nil, nil
}

// Override the Inform method for Patient
func (p *Patient) Abort() (*models.ParticipantResponse, error) {
	// Implement IDM-specific inform logic
	return nil, nil
}

func (p *Patient) Rollback() (*models.ParticipantResponse, error) {
	// Implement rollback logic
	return nil, nil
}
