package participants

import "github.com/mihnea1711/POS_Project/services/rabbit/internal/models"

// Doctor represents the Doctor module
type Doctor struct {
	models.ParticipantData // Embed Participant struct
	// Additional fields specific to Doctor, if any
}

func NewDoctor(id, name string) *Doctor {
	return &Doctor{
		ParticipantData: models.ParticipantData{
			ID:   id,
			Name: name,
		},
		// Initialize additional fields here
	}
}

// Override the Inform method for IDM
func (d *Doctor) Inform(commit bool) error {
	// Implement IDM-specific inform logic
	return nil
}

// Implement the Transactional interface methods for Participant
func (d *Doctor) Prepare() error {
	// Implement preparation logic
	return nil
}

func (d *Doctor) Commit() error {
	// Implement commit logic
	return nil
}

func (d *Doctor) Rollback() error {
	// Implement rollback logic
	return nil
}
