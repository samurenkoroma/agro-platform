package productionunit

import (
	"context"

	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func (r *productionUnitRepository) GetChildren(ctx context.Context, parentID vo.ID) ([]*pu.ProductionUnit, error) {
	query := `SELECT id, farm_id, parent_id, type, name, metadata, created_at, updated_at 
				FROM production_units
				WHERE parent_id=$1
				ORDER BY created_at`
	rows, err := r.db.Query(ctx, query, parentID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*pu.ProductionUnit, 0)

	for rows.Next() {
		var item pu.ProductionUnit

		err = rows.Scan(
			&item.ID,
			&item.FarmID,
			&item.ParentID,
			&item.Type,
			&item.Name,
			&item.Metadata,
			&item.CreatedAt,
			&item.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, &item)
	}

	return result, nil
}
