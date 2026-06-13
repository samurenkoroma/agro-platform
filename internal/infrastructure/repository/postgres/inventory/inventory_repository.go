package inventory

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	inventoryitem "github.com/samurenkoroma/agro-platform/internal/domain/inventory/aggregate/inventory_item"
	"github.com/samurenkoroma/agro-platform/internal/domain/inventory/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type inventoryRepository struct{ db uow.DB }

func NewInventoryRepository(db uow.DB) repository.InventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) Save(ctx context.Context, item *inventoryitem.Item) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO inventory_items (
			id, farm_id, name, sku, type, unit, warehouse_id,
			stock_available, stock_reserved, stock_consumed, stock_lost,
			created_at, updated_at, archived_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)
		ON CONFLICT (id) DO UPDATE SET
			name            = EXCLUDED.name,
			sku             = EXCLUDED.sku,
			warehouse_id    = EXCLUDED.warehouse_id,
			stock_available = EXCLUDED.stock_available,
			stock_reserved  = EXCLUDED.stock_reserved,
			stock_consumed  = EXCLUDED.stock_consumed,
			stock_lost      = EXCLUDED.stock_lost,
			updated_at      = EXCLUDED.updated_at,
			archived_at     = EXCLUDED.archived_at
	`,
		item.ID, item.FarmID, item.Name, item.SKU, item.Type, item.Unit, item.WarehouseID,
		item.Stock.Available, item.Stock.Reserved, item.Stock.Consumed, item.Stock.Lost,
		item.CreatedAt, item.UpdatedAt, item.ArchivedAt,
	)
	return err
}

func (r *inventoryRepository) GetByID(ctx context.Context, id vo.ID) (*inventoryitem.Item, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, farm_id, name, sku, type, unit, warehouse_id,
		       stock_available, stock_reserved, stock_consumed, stock_lost,
		       created_at, updated_at, archived_at
		FROM inventory_items WHERE id = $1`, id)
	return scanItem(row)
}

func (r *inventoryRepository) List(ctx context.Context, filter repository.ItemFilter) ([]*inventoryitem.Item, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, farm_id, name, sku, type, unit, warehouse_id,
		       stock_available, stock_reserved, stock_consumed, stock_lost,
		       created_at, updated_at, archived_at
		FROM inventory_items
		WHERE farm_id = $1 AND ($2::bool OR archived_at IS NULL)
		ORDER BY name ASC`, filter.FarmID, filter.Archived)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*inventoryitem.Item
	for rows.Next() {
		item, err := scanItem(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	return result, nil
}

func scanItem(s interface{ Scan(...any) error }) (*inventoryitem.Item, error) {
	var item inventoryitem.Item
	err := s.Scan(
		&item.ID, &item.FarmID, &item.Name, &item.SKU, &item.Type, &item.Unit, &item.WarehouseID,
		&item.Stock.Available, &item.Stock.Reserved, &item.Stock.Consumed, &item.Stock.Lost,
		&item.CreatedAt, &item.UpdatedAt, &item.ArchivedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, inventoryitem.ErrItemNotFound
		}
		return nil, err
	}
	return &item, nil
}

var _ repository.InventoryRepository = (*inventoryRepository)(nil)
