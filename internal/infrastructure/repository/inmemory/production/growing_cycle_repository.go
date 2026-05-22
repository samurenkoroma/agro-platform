package production

import (
	"context"
	"sync"

	gc "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"

	repo "github.com/samurenkoroma/agro-platform/internal/domain/production/repository"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type growingCycleRepository struct {
	mu sync.RWMutex

	items map[vo.ID]*gc.GrowingCycle
}

func (r *growingCycleRepository) Save(ctx context.Context, root *gc.GrowingCycle) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.items[root.ID] = root

	return nil
}

func (r *growingCycleRepository) GetByID(ctx context.Context, id vo.ID) (*gc.GrowingCycle, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item := r.items[id]

	return item, nil
}

func (r *growingCycleRepository) GetByFarm(ctx context.Context, farmID vo.ID) ([]*gc.GrowingCycle, error) {
	//TODO implement me
	panic("implement me")
}

func (r *growingCycleRepository) GetByProductionUnit(ctx context.Context, unitID vo.ID) ([]*gc.GrowingCycle, error) {
	//TODO implement me
	panic("implement me")
}

func (r *growingCycleRepository) Exists(ctx context.Context, id vo.ID) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r *growingCycleRepository) GetActiveByUnit(unitID vo.ID) ([]*gc.GrowingCycle, error) {
	//TODO implement me
	panic("implement me")
}

func NewGrowingCycleRepository() repo.GrowingCycleRepository {
	return &growingCycleRepository{
		items: make(map[vo.ID]*gc.GrowingCycle),
	}
}

var _ repo.GrowingCycleRepository = (*growingCycleRepository)(nil)
