package production

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	yieldbatch "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/yield_batch"
	"github.com/samurenkoroma/agro-platform/internal/domain/production/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type yieldRepository struct {
	db uow.DB
}

func (r *yieldRepository) Save(ctx context.Context, root *yieldbatch.YieldBatch) error {
	query := `
INSERT INTO yield_batches(
	id,
	growing_cycle_id,
	plant_id,
	quantity,
	fruit_count,
	grade,
	marketable,
	notes,
	harvested_at,
	metadata,
	created_at
)
VALUES(
	$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11
)
ON CONFLICT(id)
DO UPDATE SET
	quantity=excluded.quantity,
	fruit_count=excluded.fruit_count,
	grade=excluded.grade,
	marketable=excluded.marketable,
	notes=excluded.notes,
	harvested_at=excluded.harvested_at,
	metadata=excluded.metadata
`

	_, err := r.db.Exec(
		ctx,
		query,

		root.ID,

		root.GrowingCycleID,

		root.PlantID,

		root.Quantity,

		root.FruitCount,

		root.Grade,

		root.Marketable,

		root.Notes,

		root.HarvestedAt,

		root.Metadata,

		root.CreatedAt,
	)

	return err
}

func (r *yieldRepository) GetByID(ctx context.Context, id vo.ID) (*yieldbatch.YieldBatch, error) {
	query := `
SELECT
	id,
	growing_cycle_id,
	plant_id,
	quantity,
	fruit_count,
	grade,
	marketable,
	notes,
	harvested_at,
	metadata,
	created_at
FROM yield_batches
WHERE id=$1
`

	root := &yieldbatch.YieldBatch{}
	err := r.db.QueryRow(ctx, query, id).Scan(scanYieldBatch(root)...)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return root, nil
}

func (r *yieldRepository) GetByCycle(ctx context.Context, cycleID vo.ID) ([]*yieldbatch.YieldBatch, error) {

	query := `
SELECT
	id,
	growing_cycle_id,
	plant_id,
	quantity,
	fruit_count,
	grade,
	marketable,
	notes,
	harvested_at,
	metadata,
	created_at
FROM yield_batches
WHERE growing_cycle_id=$1
ORDER BY harvested_at DESC
`

	rows, err := r.db.Query(ctx, query, cycleID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*yieldbatch.YieldBatch, 0)

	for rows.Next() {
		root := &yieldbatch.YieldBatch{}
		if err := rows.Scan(scanYieldBatch(root)...); err != nil {
			return nil, err
		}

		result = append(result, root)
	}

	return result, nil
}

func (r *yieldRepository) GetByPlant(ctx context.Context, plantID vo.ID) ([]*yieldbatch.YieldBatch, error) {

	query := `SELECT
    id,growing_cycle_id,plant_id,quantity,fruit_count,grade,marketable,notes,harvested_at,metadata,created_at
		FROM yield_batches WHERE plant_id=$1 ORDER BY harvested_at DESC`

	rows, err := r.db.Query(ctx, query, plantID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*yieldbatch.YieldBatch, 0)
	for rows.Next() {
		root := &yieldbatch.YieldBatch{}

		if err := rows.Scan(scanYieldBatch(root)...); err != nil {
			return nil, err
		}

		result = append(result, root)
	}

	return result, nil
}

func (r *yieldRepository) Exists(ctx context.Context, id vo.ID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM yield_batches WHERE id=$1)`

	var exists bool
	err := r.db.QueryRow(ctx, query, id).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func NewYieldBatchRepository(db uow.DB) repository.YieldBatchRepository {
	return &yieldRepository{db: db}
}
