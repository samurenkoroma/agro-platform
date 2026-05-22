package agronomy

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/stress"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type stressRepository struct{}

func (s stressRepository) Save(ctx context.Context, root *stress.Stress) error {
	//TODO implement me
	panic("implement me")
}

func (s stressRepository) GetByID(ctx context.Context, id vo.ID) (*stress.Stress, error) {
	//TODO implement me
	panic("implement me")
}

func (s stressRepository) GetAll(ctx context.Context) ([]*stress.Stress, error) {
	//TODO implement me
	panic("implement me")
}

func (s stressRepository) Exists(ctx context.Context, id vo.ID) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewStressRepository() repository.StressRepository {
	return &stressRepository{}
}
