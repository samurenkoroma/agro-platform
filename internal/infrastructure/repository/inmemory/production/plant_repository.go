package production

import (
	"context"
	"sync"

	"github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/plant"
	"github.com/samurenkoroma/agro-platform/internal/domain/production/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type plantRepository struct {
	mu    sync.RWMutex
	items map[vo.ID]*plant.Plant
}

func (p plantRepository) Save(ctx context.Context, root *plant.Plant) error {
	//TODO implement me
	panic("implement me")
}

func (p plantRepository) GetByID(ctx context.Context, id vo.ID) (*plant.Plant, error) {
	//TODO implement me
	panic("implement me")
}

func (p plantRepository) GetByCycle(ctx context.Context, cycleID vo.ID) ([]*plant.Plant, error) {
	//TODO implement me
	panic("implement me")
}

func (p plantRepository) GetByProductionUnit(ctx context.Context, unitID vo.ID) ([]*plant.Plant, error) {
	//TODO implement me
	panic("implement me")
}

func (p plantRepository) Exists(ctx context.Context, id vo.ID) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewPlantRepository() repository.PlantRepository {
	return &plantRepository{
		items: make(map[vo.ID]*plant.Plant),
	}
}
