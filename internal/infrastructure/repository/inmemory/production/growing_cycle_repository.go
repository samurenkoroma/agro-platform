package production

import (
	"sync"

	gc "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"

	repo "github.com/samurenkoroma/agro-platform/internal/domain/production/repository"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type GrowingCycleRepository struct {
	mu sync.RWMutex

	items map[vo.ID]*gc.GrowingCycle
}

func (r *GrowingCycleRepository) GetActiveByUnit(unitID vo.ID) ([]*gc.GrowingCycle, error) {
	//TODO implement me
	panic("implement me")
}

func NewGrowingCycleRepository() *GrowingCycleRepository {
	return &GrowingCycleRepository{
		items: make(map[vo.ID]*gc.GrowingCycle),
	}
}

func (r *GrowingCycleRepository) Save(entity *gc.GrowingCycle) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.items[entity.ID] = entity

	return nil
}

func (r *GrowingCycleRepository) GetByID(id vo.ID) (*gc.GrowingCycle, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item := r.items[id]

	return item, nil
}

func (r *GrowingCycleRepository) Delete(id vo.ID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.items, id)

	return nil
}

var _ repo.GrowingCycleRepository = (*GrowingCycleRepository)(nil)
