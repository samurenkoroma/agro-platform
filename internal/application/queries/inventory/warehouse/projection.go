package warehouse

import (
	"context"
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type WarehouseDTO struct {
	ID        vo.ID     `json:"id"`
	Name      string    `json:"name"`
	Code      *string   `json:"code,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

type Projection interface {
	List(ctx context.Context, farmID vo.ID) ([]*WarehouseDTO, error)
}
