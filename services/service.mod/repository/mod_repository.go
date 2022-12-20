package repository

import (
	"context"

	"github.com/mxbikes/mxbikesclient/services/service.mod/models"
)

type ModPostgresRepository interface {
	GetModByID(ctx context.Context, modID string) (*models.Mod, error)
	SearchMod(ctx context.Context, searchText string, modTypeCategoryIDs []string, listQuery *models.ListQuery) (*models.ListResult[*models.Mod], error)
}
