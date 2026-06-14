package productionunit

import (
	"context"
	"encoding/json"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
)

func (r *productionUnitRepository) GetByID(ctx context.Context, id vo.ID) (*pu.ProductionUnit, error) {
	query := `SELECT 
    id,owner_id,parent_id,
    type,status,code,area,properties,
       created_at,updated_at
				FROM production_units
				WHERE id=$1`

	row := r.db.QueryRow(ctx, query, id)

	var root pu.ProductionUnit
	var propertiesRaw []byte
	err := row.Scan(
		&root.ID,
		&root.OwnerID,
		&root.ParentID,
		&root.Type,
		&root.Status,
		&root.Code,
		&root.Area,
		&propertiesRaw,
		&root.CreatedAt,
		&root.UpdatedAt,
	)
	if propertiesRaw != nil {
		if err := json.Unmarshal(propertiesRaw, &root.Properties); err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	return &root, nil
}
