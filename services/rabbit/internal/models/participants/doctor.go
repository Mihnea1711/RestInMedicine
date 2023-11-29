package participants

import (
	"github.com/google/uuid"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models"
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

// Override the Inform method for IDM
func (d *Doctor) Inform(commit bool) (*models.ParticipantResponse, error) {
	// Implement IDM-specific inform logic
	return nil, nil
}

// Implement the Transactional interface methods for Participant
func (d *Doctor) Prepare() (*models.ParticipantResponse, error) {
	// Implement preparation logic
	return nil, nil
}

func (d *Doctor) Commit() (*models.ParticipantResponse, error) {
	// Implement commit logic
	return nil, nil
}

func (d *Doctor) Rollback() (*models.ParticipantResponse, error) {
	// Implement rollback logic
	return nil, nil
}
