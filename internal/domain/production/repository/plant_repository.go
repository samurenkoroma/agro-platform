package repository

import (
	"context"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

	plant "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/plant"
)

type PlantRepository interface {
	Save(ctx context.Context, root *plant.Plant) error
	GetByID(ctx context.Context, id vo.ID) (*plant.Plant, error)
	GetByCycle(ctx context.Context, cycleID vo.ID) ([]*plant.Plant, error)
	GetByProductionUnit(ctx context.Context, unitID vo.ID) ([]*plant.Plant, error)
	Exists(ctx context.Context, id vo.ID) (bool, error)
}
