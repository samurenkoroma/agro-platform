package repository

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/inventory/aggregate/warehouse"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type WarehouseRepository interface {
	Save(ctx context.Context, w *warehouse.Warehouse) error
	GetByID(ctx context.Context, id vo.ID) (*warehouse.Warehouse, error)
	List(ctx context.Context, farmID vo.ID) ([]*warehouse.Warehouse, error)
}
