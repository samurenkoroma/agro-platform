package item

import (
	"context"

	itemquery "github.com/samurenkoroma/agro-platform/internal/application/queries/inventory/item"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type projection struct{ db uow.DB }

func New(db uow.DB) itemquery.Projection { return &projection{db: db} }

func (p *projection) Get(ctx context.Context, id vo.ID) (*itemquery.ItemDTO, error) {
	row := p.db.QueryRow(ctx, `
		SELECT id, name, sku, type, unit, warehouse_id,
		       stock_available, stock_reserved, stock_consumed, stock_lost, created_at
		FROM inventory_items WHERE id = $1`, id)
	return scanItemDTO(row)
}

func (p *projection) List(ctx context.Context, farmID vo.ID, warehouseID *vo.ID) ([]*itemquery.ItemDTO, error) {
	var (
		rows interface {
			Next() bool
			Scan(...any) error
			Close()
		}
		err error
	)
	if warehouseID != nil {
		rows, err = p.db.Query(ctx, `
			SELECT id, name, sku, type, unit, warehouse_id,
			       stock_available, stock_reserved, stock_consumed, stock_lost, created_at
			FROM inventory_items
			WHERE farm_id = $1 AND warehouse_id = $2 AND archived_at IS NULL
			ORDER BY name ASC`, farmID, warehouseID)
	} else {
		rows, err = p.db.Query(ctx, `
			SELECT id, name, sku, type, unit, warehouse_id,
			       stock_available, stock_reserved, stock_consumed, stock_lost, created_at
			FROM inventory_items
			WHERE farm_id = $1 AND archived_at IS NULL
			ORDER BY name ASC`, farmID)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*itemquery.ItemDTO
	for rows.Next() {
		dto, err := scanItemDTO(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, dto)
	}
	return result, nil
}

func scanItemDTO(s interface{ Scan(...any) error }) (*itemquery.ItemDTO, error) {
	var dto itemquery.ItemDTO
	return &dto, s.Scan(
		&dto.ID, &dto.Name, &dto.SKU, &dto.Type, &dto.Unit, &dto.WarehouseID,
		&dto.Stock.Available, &dto.Stock.Reserved, &dto.Stock.Consumed, &dto.Stock.Lost,
		&dto.CreatedAt,
	)
}
