package repository

import (
	"context"

	"github.com/mxbikes/mxbikesclient/services/service.modType/models"
)

type ModTypePostgresRepository interface {
	GetModTypeByID(ctx context.Context, modTypeID string) (*models.ModType, error)
	GetAllModTypes(ctx context.Context) ([]*models.ModType, error)
}
