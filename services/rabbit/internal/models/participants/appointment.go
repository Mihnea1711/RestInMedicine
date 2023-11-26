package participants

import "github.com/mihnea1711/POS_Project/services/rabbit/internal/models"

// IDM represents the IDM module
type Appointment struct {
	models.ParticipantData // Embed Participant struct
	// Additional fields specific to IDM, if any
}

func NewAppointment(id, name string) *Appointment {
	return &Appointment{
		ParticipantData: models.ParticipantData{
			ID:   id,
			Name: name,
		},
		// Initialize additional fields here
	}
}

// Override the Inform method for IDM
func (a *Appointment) Inform(commit bool) error {
	// Implement IDM-specific inform logic
	return nil
}

// Implement the Transactional interface methods for Participant
func (a *Appointment) Prepare() error {
	// Implement preparation logic
	return nil
}

func (a *Appointment) Commit() error {
	// Implement commit logic
	return nil
}

func (a *Appointment) Rollback() error {
	// Implement rollback logic
	return nil
}
