package productionunit

import (
	"context"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
)

func (r *productionUnitRepository) GetSubtree(ctx context.Context, rootID vo.ID) ([]*pu.ProductionUnit, error) {
	query := `WITH RECURSIVE tree AS (

SELECT
id,
farm_id,
parent_id,
type,
name,
metadata,
created_at,
updated_at

FROM production_units

WHERE id=$1

UNION ALL

SELECT
c.id,
c.farm_id,
c.parent_id,
c.type,
c.name,
c.metadata,
c.created_at,
c.updated_at

FROM production_units c

INNER JOIN tree t
ON c.parent_id=t.id

)

SELECT
id,
farm_id,
parent_id,
type,
name,
metadata,
created_at,
updated_at

FROM tree
ORDER BY created_at
`

	rows,
		err :=
		r.db.Query(
			ctx,
			query,
			rootID,
		)

	if err != nil {
		return nil,
			err
	}

	defer rows.Close()

	result :=
		make(
			[]*pu.
				ProductionUnit,
			0,
		)

	for rows.Next() {

		var item pu.ProductionUnit

		err =
			rows.Scan(
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
			return nil,
				err
		}

		result =
			append(
				result,
				&item,
			)
	}

	return result,
		nil
}
