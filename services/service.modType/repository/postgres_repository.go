package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mxbikes/mxbikesclient/services/service.modType/models"
)

type postgresRepository struct {
	dbPool *pgx.Conn
}

func NewPostgresRepository(c *pgx.Conn) *postgresRepository {
	return &postgresRepository{dbPool: c}
}

func (p *postgresRepository) GetModTypeByID(ctx context.Context, modTypeID string) (*models.ModType, error) {
	sqlStatement := `SELECT id, name FROM modtype WHERE id=$1`

	var modType models.ModType
	err := p.dbPool.QueryRow(ctx, sqlStatement, modTypeID).Scan(&modType.ID, &modType.Name)
	if err != nil {
		return nil, err
	}

	return &modType, nil
}

func (p *postgresRepository) GetAllModTypes(ctx context.Context) ([]*models.ModType, error) {
	sqlStatement := `SELECT id, name FROM modtype`

	rows, err := p.dbPool.Query(ctx, sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modTypes []*models.ModType
	for rows.Next() {
		var modType models.ModType
		if err := rows.Scan(&modType.ID, &modType.Name); err != nil {
			fmt.Errorf("", err)
		}

		modTypes = append(modTypes, &modType)
	}

	return modTypes, nil
}
