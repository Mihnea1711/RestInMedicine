package participants

import "github.com/mihnea1711/POS_Project/services/rabbit/internal/models"

// Patient represents the Patient module
type Patient struct {
	models.ParticipantData // Embed Participant struct
	// Additional fields specific to Patient, if any
}

func NewPatient(id, name string) *Patient {
	return &Patient{
		ParticipantData: models.ParticipantData{
			ID:   id,
			Name: name,
		},
		// Initialize additional fields here
	}
}

// Override the Inform method for IDM
func (p *Patient) Inform(commit bool) error {
	// Implement IDM-specific inform logic
	return nil
}

// Implement the Transactional interface methods for Participant
func (p *Patient) Prepare() error {
	// Implement preparation logic
	return nil
}

func (p *Patient) Commit() error {
	// Implement commit logic
	return nil
}

func (p *Patient) Rollback() error {
	// Implement rollback logic
	return nil
}
