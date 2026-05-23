package production

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/plant"
	"github.com/samurenkoroma/agro-platform/internal/domain/production/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type plantRepository struct {
	db uow.DB
}

func (r *plantRepository) Save(ctx context.Context, root *plant.Plant) error {
	query := `INSERT INTO 
    plants(id,growing_cycle_id,crop_id,variety_id,production_unit_id,slot_id,substrate_id,status,
           health,current_stage_id,planted_at,transplanted_at,harvested_at,discarded_at,metadata,created_at,updated_at)
VALUES(	$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)
ON CONFLICT(id)
DO UPDATE SET
	variety_id=excluded.variety_id,
	slot_id=excluded.slot_id,
	substrate_id=excluded.substrate_id,

	status=excluded.status,
	health=excluded.health,

	current_stage_id=excluded.current_stage_id,

	transplanted_at=excluded.transplanted_at,
	harvested_at=excluded.harvested_at,
	discarded_at=excluded.discarded_at,

	metadata=excluded.metadata,

	updated_at=excluded.updated_at
`

	_, err := r.db.Exec(
		ctx,
		query,
		root.ID,
		root.GrowingCycleID,
		root.CropID,
		root.VarietyID,
		root.ProductionUnitID,
		root.SlotID,
		root.SubstrateID,
		root.Status,
		root.Health,
		root.CurrentStageID,
		root.PlantedAt,
		root.TransplantedAt,
		root.HarvestedAt,
		root.DiscardedAt,
		root.Metadata,
		root.CreatedAt,
		root.UpdatedAt,
	)

	return err
}

func (r *plantRepository) GetByID(ctx context.Context, id vo.ID) (*plant.Plant, error) {
	query := `SELECT
	id,growing_cycle_id,crop_id,variety_id,production_unit_id,slot_id,substrate_id,status,health,
	current_stage_id,planted_at,transplanted_at,harvested_at,discarded_at,metadata,created_at,updated_at
FROM plants WHERE id=$1`

	root := &plant.Plant{}

	err := r.db.QueryRow(ctx, query, id).Scan(scanPlant(root)...)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return root, nil
}

func (r *plantRepository) GetByCycle(ctx context.Context, cycleID vo.ID) ([]*plant.Plant, error) {
	query := `SELECT id,growing_cycle_id,crop_id,variety_id,production_unit_id,slot_id,substrate_id,status,health,
       current_stage_id,planted_at,transplanted_at,harvested_at,discarded_at,metadata,created_at,updated_at
FROM plants WHERE growing_cycle_id=$1 ORDER BY planted_at`

	rows, err := r.db.Query(ctx, query, cycleID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*plant.Plant, 0)

	for rows.Next() {
		root := &plant.Plant{}

		if err := rows.Scan(scanPlant(root)...); err != nil {
			return nil, err
		}

		result = append(result, root)
	}

	return result, nil
}

func (r *plantRepository) GetByProductionUnit(ctx context.Context, unitID vo.ID) ([]*plant.Plant, error) {
	query := `SELECT id,growing_cycle_id,crop_id,variety_id,production_unit_id,slot_id,substrate_id,status,health,
       current_stage_id,planted_at,transplanted_at,harvested_at,discarded_at,metadata,created_at,updated_at
FROM plants WHERE production_unit_id=$1 ORDER BY planted_at`

	rows, err := r.db.Query(ctx, query, unitID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*plant.Plant, 0)

	for rows.Next() {
		root := &plant.Plant{}

		if err := rows.Scan(scanPlant(root)...); err != nil {
			return nil, err
		}

		result = append(result, root)
	}

	return result, nil
}

func (r *plantRepository) Exists(ctx context.Context, id vo.ID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM plants WHERE id=$1)`

	var exists bool

	err := r.db.QueryRow(ctx, query, id).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func NewPlantRepository(db uow.DB) repository.PlantRepository {
	return &plantRepository{db: db}
}
