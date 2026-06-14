package warehouse

import (
	"context"

	whquery "github.com/samurenkoroma/agro-platform/internal/application/queries/inventory/warehouse"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type projection struct{ db uow.DB }

func New(db uow.DB) whquery.Projection { return &projection{db: db} }

func (p *projection) List(ctx context.Context, farmID vo.ID) ([]*whquery.WarehouseDTO, error) {
	rows, err := p.db.Query(ctx, `
		SELECT id, name, code, created_at
		FROM inventory_warehouses
		WHERE farm_id = $1 AND archived_at IS NULL
		ORDER BY name ASC`, farmID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*whquery.WarehouseDTO, 0)
	for rows.Next() {
		var dto whquery.WarehouseDTO
		if err := rows.Scan(&dto.ID, &dto.Name, &dto.Code, &dto.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, &dto)
	}
	return result, nil
}
