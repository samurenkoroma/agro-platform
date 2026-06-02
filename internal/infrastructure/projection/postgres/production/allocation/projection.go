package allocation

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/queries/production/allocation"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type projection struct {
	db uow.DB
}

func New(db uow.DB) allocation.Projection {
	return &projection{db: db}
}

func (p *projection) Get(ctx context.Context, id vo.ID) (*allocation.DTO, error) {
	sql := `SELECT
    a.id,
    a.cycle_id,
    a.production_unit_id,
    unit.code,
    a.area,
    a.started_at,
    a.ended_at
FROM public.production_allocations a
         left join production_units unit on a.production_unit_id = a.id
WHERE a.id = $1`

	row := p.db.QueryRow(ctx, sql, id)

	return scanDTO(row)
}

func (p *projection) List(ctx context.Context, ownerId vo.ID) ([]*allocation.DTO, error) {
	sql := `SELECT
    a.id,
    a.cycle_id,
    a.production_unit_id,
    unit.code,
    a.area,
    a.started_at,
    a.ended_at
FROM public.production_allocations a
         left join production_units unit on a.production_unit_id = a.id
WHERE unit.owner_id = $1  ORDER BY unit.code`
	rows, err := p.db.Query(ctx, sql, ownerId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*allocation.DTO, 0)

	for rows.Next() {
		item, err := scanDTO(rows)

		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	return result, nil
}

type scanner interface {
	Scan(dest ...any) error
}

func scanDTO(row scanner) (*allocation.DTO, error) {
	var result allocation.DTO

	if err := row.Scan(
		&result.ID, &result.CycleID, &result.ProductionUnitID, &result.ProductionUnitName,
		&result.Area, &result.StartedAt, &result.EndedAt,
	); err != nil {
		return nil, err
	}

	return &result, nil
}
