package productionunit

import (
	"context"

	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
)

func (r *productionUnitRepository) Save(ctx context.Context, unit *pu.ProductionUnit) error {
	query := `INSERT INTO 
    production_units(
                     id,owner_id,parent_id, code,
                     status,type,properties,
                     created_at,updated_at
                     ) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)
				ON CONFLICT(id) 
				DO UPDATE SET
				    parent_id=excluded.parent_id,
					updated_at=excluded.updated_at`

	_, err := r.db.Exec(ctx, query,
		unit.ID, unit.OwnerID, unit.ParentID, unit.Code,
		unit.Status, unit.Type, unit.Properties,
		unit.CreatedAt, unit.UpdatedAt,
	)

	return err
}
