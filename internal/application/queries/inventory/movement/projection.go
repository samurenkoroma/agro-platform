package movement

import (
	"context"
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type MovementDTO struct {
	ID              vo.ID     `json:"id"`
	ItemID          vo.ID     `json:"itemId"`
	Type            string    `json:"type"`
	Quantity        float64   `json:"quantity"`
	FromWarehouseID *vo.ID    `json:"fromWarehouseId,omitempty"`
	ToWarehouseID   *vo.ID    `json:"toWarehouseId,omitempty"`
	ReferenceType   *string   `json:"referenceType,omitempty"`
	ReferenceID     *string   `json:"referenceId,omitempty"`
	Note            *string   `json:"note,omitempty"`
	Timestamp       time.Time `json:"timestamp"`
}

type Projection interface {
	List(ctx context.Context, farmID vo.ID, itemID *vo.ID) ([]*MovementDTO, error)
}
