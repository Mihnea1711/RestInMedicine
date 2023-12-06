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
	return &Doctor{
		Participant: *models.NewParticipant(participantID, participantType),
	}
}

// Implement the Transactional interface methods for Participant
func (d *Doctor) Prepare() (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending doctor prepare request")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), utils.REQUEST_TIMEOUT_MULTIPLIER*time.Second)
	defer cancel() // Make sure to call cancel to release resources associated with the context

	response, status, err := utils.MakeRequest(ctx, http.MethodGet, utils.DOCTOR_HOST, utils.PREPARE_DOCTOR_ENDPOINT, utils.DOCTOR_PORT, nil)
	if err != nil {
		log.Printf("Error making doctor prepare request: %v", err)
		return nil, err
	}

	return &models.ParticipantResponse{
		Code:    status,
		Message: response.Message,
	}, nil
}

func (d *Doctor) Commit(userID int) (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending doctor commit request")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), utils.REQUEST_TIMEOUT_MULTIPLIER*time.Second)
	defer cancel()

	response, status, err := utils.MakeRequest(ctx, http.MethodPatch, utils.DOCTOR_HOST, utils.COMMIT_DOCTOR_ENDPOINT, utils.DOCTOR_PORT, models.ActivityData{
		IsActive: false,
		IDUser:   userID,
	})
	if err != nil {
		log.Printf("Error making doctor prepare request: %v", err)
		return nil, err
	}

	return &models.ParticipantResponse{
		Code:    status,
		Message: response.Message,
	}, nil
}

func (d *Doctor) Abort() (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending doctor abort request")

	return &models.ParticipantResponse{
		Code:    http.StatusOK,
		Message: "Transaction aborted successfully",
	}, nil
}

func (d *Doctor) Rollback() (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending doctor rollback request")

	return &models.ParticipantResponse{
		Code:    http.StatusOK,
		Message: "Transaction rolled back successfully",
	}, nil
}

func (d *Doctor) Compensate() error {
	return nil
}
