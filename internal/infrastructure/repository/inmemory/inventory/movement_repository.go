package inventory

import (
	"context"
	"sync"

	"github.com/samurenkoroma/agro-platform/internal/domain/inventory/aggregate/movement"
	"github.com/samurenkoroma/agro-platform/internal/domain/inventory/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type movementRepository struct {
	mu    sync.RWMutex
	items map[vo.ID]*movement.Movement
}

func NewMovementRepository() repository.MovementRepository {
	return &movementRepository{items: make(map[vo.ID]*movement.Movement)}
}

func (r *movementRepository) Save(ctx context.Context, m *movement.Movement) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[m.ID] = m
	return nil
}

func (r *movementRepository) GetByID(ctx context.Context, id vo.ID) (*movement.Movement, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	m, ok := r.items[id]
	if !ok {
		return nil, movement.ErrInvalidQuantity
	}
	return m, nil
}

func (r *movementRepository) List(ctx context.Context, filter repository.MovementFilter) ([]*movement.Movement, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*movement.Movement
	for _, m := range r.items {
		if m.FarmID != filter.FarmID {
			continue
		}
		if filter.ItemID != nil && m.ItemID != *filter.ItemID {
			continue
		}
		if filter.Type != nil && m.Type != *filter.Type {
			continue
		}
		result = append(result, m)
	}
	return result, nil
}

var _ repository.MovementRepository = (*movementRepository)(nil)
