package mysql

import (
	"context"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
)

func (db *MySQLDatabase) UpdatePacientByID(ctx context.Context, pacient *models.Pacient) (int64, error) {
	// TODO
	return 0, nil
}
