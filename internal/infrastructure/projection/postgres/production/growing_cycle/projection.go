package growingcycle

import (
	"context"

	growingcycle "github.com/samurenkoroma/agro-platform/internal/application/queries/production/growing_cycle"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type projection struct {
	db uow.DB
}

func New(db uow.DB) growingcycle.Projection {
	return &projection{db: db}
}

func (p *projection) Get(ctx context.Context, id vo.ID) (*growingcycle.DTO, error) {
	sql := `SELECT cycle.id,
       cycle.name,
       code,
       c.name crop_aame,
       v.name variety_name,
       status,
       stage,
       expected_harvest_at,
       cycle.created_at
FROM production_growing_cycles cycle
         left join crops c on crop_id = c.id
         left join varieties v on variety_id = v.id  WHERE cycle.id = $1`

	row := p.db.QueryRow(ctx, sql, id)

	return scanDTO(row)
}

func (p *projection) List(ctx context.Context, ownerId vo.ID) ([]*growingcycle.DTO, error) {
	sql := `SELECT cycle.id,
       cycle.name,
       code,
       c.name crop_aame,
       v.name variety_name,
       status,
       stage,
       expected_harvest_at,
       cycle.created_at
FROM production_growing_cycles cycle
         left join crops c on crop_id = c.id
         left join varieties v on variety_id = v.id  WHERE farm_id = $1  ORDER BY code`
	rows, err := p.db.Query(ctx, sql, ownerId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*growingcycle.DTO, 0)

	for rows.Next() {
		item, err := scanDTO(rows)

		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	return result, nil
}

func (p *projection) Summary(ctx context.Context, ownerId vo.ID, cycleId vo.ID) (*growingcycle.SummaryDTO, error) {
	sql := `SELECT
    cycle.id,
    cycle.name,
    cycle.status,

    COALESCE((
        SELECT SUM(area)
        FROM production_allocations
        WHERE cycle_id = cycle.id
          AND ended_at IS NULL
    ),0) allocated_area,

    COALESCE((
        SELECT SUM(quantity)
        FROM production_plantings
        WHERE cycle_id = cycle.id
    ),0) planted_quantity,

    COALESCE((
        SELECT SUM(quantity)
        FROM public.production_harvest_batch
        WHERE cycle_id = cycle.id
    ),0) harvested_quantity
FROM production_growing_cycles cycle WHERE  cycle.farm_id = $1 and cycle.id = $2`
	row := p.db.QueryRow(ctx, sql, ownerId, cycleId)

	var sum growingcycle.SummaryDTO
	if err := row.Scan(&sum.ID, &sum.Name, &sum.Status, &sum.AllocatedArea, &sum.PlantedQuantity, &sum.HarvestedQuantity); err != nil {
		return nil, err
	}

	return &sum, nil
}

type scanner interface {
	Scan(dest ...any) error
}

func scanDTO(row scanner) (*growingcycle.DTO, error) {
	var result growingcycle.DTO

	if err := row.Scan(
		&result.ID, &result.Name, &result.Code,
		&result.CropName, &result.VarietyName,
		&result.Status, &result.Stage,
		&result.ExpectedHarvestAt, &result.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &result, nil
}
