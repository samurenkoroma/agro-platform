package productionunit

import (
	"context"
	"encoding/json"
	"fmt"

	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
)

func (r *productionUnitRepository) Save(ctx context.Context, unit *pu.ProductionUnit) error {
	query := `INSERT INTO 
    production_units(
                     id,owner_id,parent_id, code, area,
                     status,type,properties,
                     created_at,updated_at
                     ) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
				ON CONFLICT(id) 
				DO UPDATE SET
				    parent_id=excluded.parent_id,
					updated_at=excluded.updated_at,
					status=excluded.status`

	var propsJSON []byte
	var err error
	props := unit.Properties
	propsJSON, err = json.Marshal(props)
	if err != nil {
		return fmt.Errorf("failed to marshal bed attributes: %w", err)
	}

	_, err = r.db.Exec(ctx, query,
		unit.ID, unit.OwnerID, unit.ParentID, unit.Code, unit.Area,
		unit.Status, unit.Type, propsJSON,
		unit.CreatedAt, unit.UpdatedAt,
	)

	return err
}
