package productionunit

import (
	"context"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
)

func (r *productionUnitRepository) GetByID(ctx context.Context, id vo.ID) (*pu.ProductionUnit, error) {
	query := `SELECT id,owner_id,parent_id,type,created_at,updated_at
				FROM production_units
				WHERE id=$1`

	row := r.db.QueryRow(ctx, query, id)

	var root pu.ProductionUnit

	err := row.Scan(
		&root.ID,
		&root.OwnerID,
		&root.ParentID,
		&root.Type,
		&root.CreatedAt,
		&root.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &root, nil
}
