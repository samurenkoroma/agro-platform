package repository

import (
	"context"

	gc "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type GrowingCycleRepository interface {
	Save(ctx context.Context, root *gc.GrowingCycle) error
	GetByID(ctx context.Context, id vo.ID) (*gc.GrowingCycle, error)
	GetByFarm(ctx context.Context, farmID vo.ID) ([]*gc.GrowingCycle, error)
	GetByProductionUnit(ctx context.Context, unitID vo.ID) ([]*gc.GrowingCycle, error)
	Exists(ctx context.Context, id vo.ID) (bool, error)
}
