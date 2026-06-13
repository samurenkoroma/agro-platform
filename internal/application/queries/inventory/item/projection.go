package item

import (
	"context"
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type StockDTO struct {
	Available float64 `json:"available"`
	Reserved  float64 `json:"reserved"`
	Consumed  float64 `json:"consumed"`
	Lost      float64 `json:"lost"`
}

type ItemDTO struct {
	ID          vo.ID     `json:"id"`
	Name        string    `json:"name"`
	SKU         *string   `json:"sku,omitempty"`
	Type        string    `json:"type"`
	Unit        string    `json:"unit"`
	WarehouseID *vo.ID    `json:"warehouseId,omitempty"`
	Stock       StockDTO  `json:"stock"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Projection interface {
	Get(ctx context.Context, id vo.ID) (*ItemDTO, error)
	List(ctx context.Context, farmID vo.ID, warehouseID *vo.ID) ([]*ItemDTO, error)
}
