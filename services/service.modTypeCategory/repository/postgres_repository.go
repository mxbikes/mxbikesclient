package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mxbikes/mxbikesclient/services/service.modTypeCategory/models"
)

type postgresRepository struct {
	dbPool *pgxpool.Pool
}

func NewPostgresRepository(c *pgxpool.Pool) *postgresRepository {
	return &postgresRepository{dbPool: c}
}

func (p *postgresRepository) GetModTypeCategoryByID(ctx context.Context, modTypeCategoryID string) (*models.ModTypeCategory, error) {
	sqlStatement := `SELECT id, name, id_modtype FROM modtypecategory WHERE id=$1`

	var modTypeCategory models.ModTypeCategory
	err := p.dbPool.QueryRow(ctx, sqlStatement, modTypeCategoryID).Scan(&modTypeCategory.ID, &modTypeCategory.Name, &modTypeCategory.ModTypeID)
	if err != nil {
		return nil, err
	}

	return &modTypeCategory, nil
}

func (p *postgresRepository) GetModTypeCategoriesByModTypeID(ctx context.Context, modTypeID string) ([]*models.ModTypeCategory, error) {
	sqlStatement := `SELECT id, name, id_modtype FROM modtypecategory WHERE id_modtype=$1`

	rows, err := p.dbPool.Query(ctx, sqlStatement, modTypeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modTypeCategories []*models.ModTypeCategory
	for rows.Next() {
		var modTypeCategory models.ModTypeCategory
		if err := rows.Scan(&modTypeCategory.ID, &modTypeCategory.Name, &modTypeCategory.ModTypeID); err != nil {
			fmt.Errorf("", err)
		}

		modTypeCategories = append(modTypeCategories, &modTypeCategory)
	}

	return modTypeCategories, nil
}

func (p *postgresRepository) GetAllModTypeCategories(ctx context.Context) ([]*models.ModTypeCategory, error) {
	sqlStatement := `SELECT id, name, id_modtype FROM modtypecategory`

	rows, err := p.dbPool.Query(ctx, sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modTypeCategories []*models.ModTypeCategory
	for rows.Next() {
		var modTypeCategory models.ModTypeCategory
		if err := rows.Scan(&modTypeCategory.ID, &modTypeCategory.Name, &modTypeCategory.ModTypeID); err != nil {
			fmt.Errorf("", err)
		}

		modTypeCategories = append(modTypeCategories, &modTypeCategory)
	}

	return modTypeCategories, nil
}
