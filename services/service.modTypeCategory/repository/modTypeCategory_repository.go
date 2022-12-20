package repository

import (
	"context"

	"github.com/mxbikes/mxbikesclient/services/service.modTypeCategory/models"
)

type ModTypeCategoryPostgresRepository interface {
	GetModTypeCategoryByID(ctx context.Context, modTypeCategoryID string) (*models.ModTypeCategory, error)
	GetModTypeCategoriesByModTypeID(ctx context.Context, modTypeID string) ([]*models.ModTypeCategory, error)
	GetAllModTypeCategories(ctx context.Context) ([]*models.ModTypeCategory, error)
}
