package production

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	harvest "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/harvest_batch"
	"github.com/samurenkoroma/agro-platform/internal/domain/production/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type harvestBatchRepository struct {
	db uow.DB
}

func (r *harvestBatchRepository) Save(ctx context.Context, root *harvest.HarvestBatch) error {
	query := `INSERT INTO 
    harvest_batches(id,growing_cycle_id,production_unit_id,quantity,harvested_area,grade,marketable,notes,harvested_at,metadata,created_at)
		VALUES(	$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		ON CONFLICT(id)
		DO UPDATE SET
	quantity=excluded.quantity,
	harvested_area=excluded.harvested_area,
	grade=excluded.grade,
	marketable=excluded.marketable,
	notes=excluded.notes,
	harvested_at=excluded.harvested_at,
	metadata=excluded.metadata`

	_, err := r.db.Exec(
		ctx,
		query,

		root.ID,

		root.GrowingCycleID,
		root.ProductionUnitID,

		root.Quantity,

		root.HarvestedArea,

		root.Grade,

		root.Marketable,

		root.Notes,

		root.HarvestedAt,

		root.Metadata,

		root.CreatedAt,
	)

	return err
}

func (r *harvestBatchRepository) GetByID(ctx context.Context, id vo.ID) (*harvest.HarvestBatch, error) {
	query := `SELECT
    id,growing_cycle_id,production_unit_id,quantity,harvested_area,grade,marketable,notes,harvested_at,metadata,created_at
FROM harvest_batches WHERE id=$1`

	root := &harvest.HarvestBatch{}
	err := r.db.QueryRow(ctx, query, id).Scan(scanHarvestBatch(root)...)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return root, nil
}

func (r *harvestBatchRepository) GetByCycle(ctx context.Context, cycleID vo.ID) ([]*harvest.HarvestBatch, error) {
	query := `SELECT
	id,growing_cycle_id,production_unit_id,quantity,harvested_area,grade,marketable,notes,harvested_at,metadata,created_at
FROM harvest_batches WHERE growing_cycle_id=$1 ORDER BY harvested_at DESC`

	rows, err := r.db.Query(ctx, query, cycleID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*harvest.HarvestBatch, 0)

	for rows.Next() {
		root := &harvest.HarvestBatch{}

		if err := rows.Scan(scanHarvestBatch(root)...); err != nil {
			return nil, err
		}

		result = append(result, root)
	}

	return result, nil
}

func (r *harvestBatchRepository) GetByProductionUnit(ctx context.Context, unitID vo.ID) ([]*harvest.HarvestBatch, error) {
	query := `SELECT
id,growing_cycle_id,production_unit_id,quantity,harvested_area,grade,marketable,notes,harvested_at,metadata,created_at
FROM harvest_batches WHERE production_unit_id=$1 ORDER BY harvested_at DESC`

	rows, err := r.db.Query(ctx, query, unitID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*harvest.HarvestBatch, 0)
	for rows.Next() {
		root := &harvest.HarvestBatch{}

		if err := rows.Scan(scanHarvestBatch(root)...); err != nil {
			return nil, err
		}

		result = append(result, root)
	}

	return result, nil
}

func (r *harvestBatchRepository) Exists(ctx context.Context, id vo.ID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM harvest_batches WHERE id=$1)`

	var exists bool

	err := r.db.QueryRow(ctx, query, id).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func NewHarvestBatchRepository(db uow.DB) repository.HarvestBatchRepository {
	return &harvestBatchRepository{
		db: db,
	}
}
