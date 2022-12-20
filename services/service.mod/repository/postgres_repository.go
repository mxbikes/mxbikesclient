package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	"github.com/mxbikes/mxbikesclient/services/service.mod/models"
	"golang.org/x/exp/slices"
)

type postgresRepository struct {
	dbPool *pgxpool.Pool
}

func NewPostgresRepository(c *pgxpool.Pool) *postgresRepository {
	return &postgresRepository{dbPool: c}
}

func (p *postgresRepository) GetModByID(ctx context.Context, modID string) (*models.Mod, error) {
	sqlStatement := `SELECT id, name, description, id_modtypecategory, releaseyear, created_at FROM mod WHERE id=$1`

	var mod models.Mod
	err := p.dbPool.QueryRow(ctx, sqlStatement, modID).Scan(&mod.ID, &mod.Name, &mod.Description, &mod.ModTypeCategoryID, &mod.ReleaseYear, &mod.CreateAt)
	if err != nil {
		return nil, err
	}

	return &mod, nil
}

func (p *postgresRepository) SearchMod(ctx context.Context, searchText string, modTypeCategoryIDs []string, listQuery *models.ListQuery) (*models.ListResult[*models.Mod], error) {
	limit := listQuery.GetLimit()
	skip := listQuery.GetOffset()

	if len(modTypeCategoryIDs) <= 0 {
		modTypeCategoryIDs = append(modTypeCategoryIDs, "all")
	}

	var count int64
	sqlStatement := `SELECT COUNT(*) FROM mod WHERE LOWER(name) LIKE (LOWER($1)) AND (id_modtypecategory::text=ANY($2))`
	if slices.Contains(modTypeCategoryIDs, "all") {
		modTypeCategoryIDs = append(modTypeCategoryIDs, "empty")
		sqlStatement = `SELECT COUNT(*) FROM mod WHERE LOWER(name) LIKE (LOWER($1)) AND NOT (id_modtypecategory::text=ANY($2))`
	}

	rows, err := p.dbPool.Query(ctx, sqlStatement, searchText, pq.StringArray(modTypeCategoryIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Scan(&count)

	sqlStatement = `SELECT id, name, description, id_modtypecategory, releaseyear, created_at FROM mod WHERE LOWER(name) LIKE (LOWER($1)) AND (id_modtypecategory::text=ANY($2)) LIMIT $3 OFFSET $4`
	if slices.Contains(modTypeCategoryIDs, "all") {
		fmt.Print("EMTPY")
		sqlStatement = `SELECT id, name, description, id_modtypecategory, releaseyear, created_at FROM mod WHERE LOWER(name) LIKE (LOWER($1)) AND NOT (id_modtypecategory::text=ANY($2)) LIMIT $3 OFFSET $4`
	}
	rows, err = p.dbPool.Query(ctx, sqlStatement, searchText, pq.StringArray(modTypeCategoryIDs), limit, skip)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mods := make([]*models.Mod, 0, listQuery.GetSize())

	for rows.Next() {
		var mod models.Mod
		if err := rows.Scan(&mod.ID, &mod.Name, &mod.Description, &mod.ModTypeCategoryID, &mod.ReleaseYear, &mod.CreateAt); err != nil {
			fmt.Errorf("", err)
		}

		mods = append(mods, &mod)
	}

	return models.NewListResult(mods, listQuery.GetSize(), listQuery.GetPage(), count), nil
}
