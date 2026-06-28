package productionunit

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
	spatial "github.com/samurenkoroma/agro-platform/internal/domain/spatial/repository"
)

type productionUnitRepository struct {
	db uow.DB
}

func (r *productionUnitRepository) GetNextSequence(ctx context.Context, orgID vo.ID, parentID *vo.ID, unitType pu.ProductionUnitType) (int, error) {
	sql := `SELECT COALESCE(MAX(sequence),0)+1 
				FROM production_units
			WHERE owner_id = $1
			  AND type = $2
			  AND parent_id IS NOT DISTINCT FROM $3
`

	var next int
	err := r.db.QueryRow(ctx, sql, orgID, unitType, parentID).Scan(&next)

	return next, err
}

func (r *productionUnitRepository) GetParentCode(ctx context.Context, parentID vo.ID) (string, error) {
	sql := `SELECT code FROM production_units WHERE id = $1`
	var code string
	err := r.db.QueryRow(ctx, sql, parentID).Scan(&code)
	return code, err
}

func New(db uow.DB) spatial.ProductionUnitRepository {
	return &productionUnitRepository{db: db}
}
