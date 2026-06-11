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

	sql := `
SELECT
    cycle.id,
    crop.name,
    variety.name,

    cycle.status,
    cycle.stage,

    COALESCE(SUM(a.area), 0) allocated_area,

    NULL::integer tasks_count,
    0 progress,

    MIN(a.started_at) start_date,
    MAX(a.ended_at) end_date

FROM production_growing_cycles cycle

         INNER JOIN crops crop
                    ON crop.id = cycle.crop_id

         LEFT JOIN varieties variety
                   ON variety.id = cycle.variety_id

         LEFT JOIN production_allocations a
                   ON a.cycle_id = cycle.id

WHERE cycle.farm_id = $1

GROUP BY
    cycle.id,
    crop.name,
    variety.name,
    cycle.status,
    cycle.stage,
    cycle.created_at

ORDER BY cycle.created_at DESC
`

	rows, err := p.db.Query(ctx, sql, ownerId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*growingcycle.DTO, 0)
	cycleIDs := make([]vo.ID, 0)

	index := make(map[vo.ID]*growingcycle.DTO)

	for rows.Next() {

		item := &growingcycle.DTO{}

		if err := rows.Scan(
			&item.ID,
			&item.CropName,
			&item.VarietyName,

			&item.Status,
			&item.Stage,

			&item.AllocatedArea,

			&item.TasksCount,
			&item.Progress,

			&item.StartDate,
			&item.EndDate,
		); err != nil {
			return nil, err
		}

		item.Allocations = make([]growingcycle.AllocationDTO, 0)

		result = append(result, item)
		cycleIDs = append(cycleIDs, item.ID)

		index[item.ID] = item
	}

	if len(cycleIDs) == 0 {
		return result, nil
	}

	sql = `
SELECT
    a.cycle_id,

    a.production_unit_id,
    pu.code,

    a.area,

    0 progress,

    a.started_at,
    a.ended_at

FROM production_allocations a

         INNER JOIN production_units pu
                    ON pu.id = a.production_unit_id

WHERE a.cycle_id = ANY($1)
ORDER BY pu.code
`

	rows, err = p.db.Query(ctx, sql, cycleIDs)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var cycleID vo.ID

		var allocation growingcycle.AllocationDTO

		if err := rows.Scan(
			&cycleID,

			&allocation.ProductionUnitId,
			&allocation.ProductionUnitName,

			&allocation.Area,

			&allocation.Progress,

			&allocation.StartDate,
			&allocation.EndDate,
		); err != nil {
			return nil, err
		}

		dto, ok := index[cycleID]

		if !ok {
			continue
		}

		dto.Allocations = append(dto.Allocations, allocation)
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
	var item growingcycle.DTO

	if err := row.Scan(
		&item.ID,
		&item.CropName,
		&item.VarietyName,

		&item.Status,
		&item.Stage,

		&item.AllocatedArea,

		&item.TasksCount,
		&item.Progress,

		&item.StartDate,
		&item.EndDate,
	); err != nil {
		return nil, err
	}

	return &item, nil
}
