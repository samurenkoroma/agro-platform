package repository

import (
	"context"

	inventoryitem "github.com/samurenkoroma/agro-platform/internal/domain/inventory/aggregate/inventory_item"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type ItemFilter struct {
	FarmID      vo.ID
	WarehouseID *vo.ID
	Type        *inventoryitem.Type
	Archived    bool
}

type InventoryRepository interface {
	Save(ctx context.Context, item *inventoryitem.Item) error
	GetByID(ctx context.Context, id vo.ID) (*inventoryitem.Item, error)
	List(ctx context.Context, filter ItemFilter) ([]*inventoryitem.Item, error)
}
