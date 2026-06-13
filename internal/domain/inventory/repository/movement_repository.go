package repository

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/inventory/aggregate/movement"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type MovementFilter struct {
	FarmID vo.ID
	ItemID *vo.ID
	Type   *movement.Type
}

type MovementRepository interface {
	Save(ctx context.Context, m *movement.Movement) error
	GetByID(ctx context.Context, id vo.ID) (*movement.Movement, error)
	List(ctx context.Context, filter MovementFilter) ([]*movement.Movement, error)
}
