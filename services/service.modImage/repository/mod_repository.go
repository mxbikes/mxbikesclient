package repository

import (
	"context"

	"github.com/mxbikes/mxbikesclient/services/service.modImage/models"
)

type ModImageMinioRepository interface {
	GetModImagesByModID(ctx context.Context, modID string) ([]*models.ModImage, error)
}
