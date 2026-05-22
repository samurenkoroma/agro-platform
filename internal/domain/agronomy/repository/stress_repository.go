package repository

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/stress"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type StressRepository interface {
	Save(ctx context.Context, root *stress.Stress) error
	GetByID(ctx context.Context, id vo.ID) (*stress.Stress, error)
	GetAll(ctx context.Context) ([]*stress.Stress, error)
	Exists(ctx context.Context, id vo.ID) (bool, error)
}
