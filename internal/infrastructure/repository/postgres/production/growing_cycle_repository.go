package growingcycle

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	gc "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"
	"github.com/samurenkoroma/agro-platform/internal/domain/production/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type growingCycleRepository struct {
	db uow.DB
}

func (r *growingCycleRepository) Save(ctx context.Context, root *gc.GrowingCycle) error {
	query := `INSERT INTO
    growing_cycles(id,farm_id,crop_id,production_unit_id,method,status,metadata,created_at,updated_at,archived_at)
		VALUES(	$1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
		ON CONFLICT(id)
		DO UPDATE SET
			status=excluded.status,
			method=excluded.method,
			metadata=excluded.metadata,
			updated_at=excluded.updated_at,
			archived_at=excluded.archived_at`

	_, err := r.db.Exec(ctx, query,
		root.ID,

		root.FarmID,
		root.CropID,

		root.ProductionUnitID,

		root.Method,
		root.Status,

		root.Metadata,

		root.CreatedAt,
		root.UpdatedAt,
		root.ArchivedAt,
	)

	return err
}

func (r *growingCycleRepository) GetByID(ctx context.Context, id vo.ID) (*gc.GrowingCycle, error) {
	query := `SELECT id,farm_id,crop_id,production_unit_id,method,status,metadata,created_at,updated_at,archived_at
				FROM growing_cycles
				WHERE id=$1`

	root := &gc.GrowingCycle{}

	err := r.db.QueryRow(ctx, query, id).Scan(scanGrowingCycle(root)...)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return root, nil
}

func (r *growingCycleRepository) GetByFarm(ctx context.Context, farmID vo.ID) ([]*gc.GrowingCycle, error) {
	query := `SELECT id,farm_id,crop_id,production_unit_id,method,status,metadata,created_at,updated_at,archived_at
				FROM growing_cycles
				WHERE farm_id=$1
				ORDER BY created_at DESC`

	rows, err := r.db.Query(ctx, query, farmID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*gc.GrowingCycle, 0)

	for rows.Next() {
		root := &gc.GrowingCycle{}

		if err := rows.Scan(scanGrowingCycle(root)...); err != nil {
			return nil, err
		}

		result = append(result, root)
	}

	return result, nil
}

func (r *growingCycleRepository) GetByProductionUnit(ctx context.Context, unitID vo.ID) ([]*gc.GrowingCycle, error) {
	query := `SELECT id,farm_id,crop_id,production_unit_id,method,status,metadata,created_at,updated_at,archived_at
				FROM growing_cycles
				WHERE production_unit_id=$1
				ORDER BY created_at DESC`

	rows, err := r.db.Query(ctx, query, unitID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*gc.GrowingCycle, 0)

	for rows.Next() {
		root := &gc.GrowingCycle{}

		if err := rows.Scan(scanGrowingCycle(root)...); err != nil {
			return nil, err
		}

		result = append(result, root)
	}

	return result, nil
}

func (r *growingCycleRepository) Exists(ctx context.Context, id vo.ID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM growing_cycles WHERE id=$1)`

	var exists bool

	err := r.db.QueryRow(ctx, query, id).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func NewGrowingCycleRepository(db uow.DB) repository.GrowingCycleRepository {
	return &growingCycleRepository{
		db: db,
	}
}
