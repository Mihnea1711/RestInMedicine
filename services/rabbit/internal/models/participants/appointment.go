package participants

import (
	"github.com/google/uuid"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models"
)

// IDM represents the IDM module
type Appointment struct {
	models.Participant // Embed Participant struct
	// Additional fields specific to IDM, if any
}

func NewAppointment(participantID uuid.UUID, participantType models.ParticipantType) *Appointment {
	return &Appointment{
		Participant: *models.NewParticipant(participantID, participantType),
	}
}

// Override the Inform method for IDM
func (a *Appointment) Inform(commit bool) (*models.ParticipantResponse, error) {
	// Implement IDM-specific inform logic
	return nil, nil
}

// Implement the Transactional interface methods for Participant
func (a *Appointment) Prepare() (*models.ParticipantResponse, error) {
	// Implement preparation logic
	return nil, nil
}

func (a *Appointment) Commit() (*models.ParticipantResponse, error) {
	// Implement commit logic
	return nil, nil
}

func (a *Appointment) Rollback() (*models.ParticipantResponse, error) {
	// Implement rollback logic
	return nil, nil
}
