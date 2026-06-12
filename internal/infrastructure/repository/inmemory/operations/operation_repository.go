package operations

import (
	"context"
	"sync"

	operationevent "github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/operation_event"
	"github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type operationRepository struct {
	mu    sync.RWMutex
	items map[vo.ID]*operationevent.OperationEvent
}

func NewOperationRepository() repository.OperationRepository {
	return &operationRepository{items: make(map[vo.ID]*operationevent.OperationEvent)}
}

func (r *operationRepository) Save(ctx context.Context, e *operationevent.OperationEvent) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[e.ID] = e
	return nil
}

func (r *operationRepository) GetByID(ctx context.Context, id vo.ID) (*operationevent.OperationEvent, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	e, ok := r.items[id]
	if !ok {
		return nil, operationevent.ErrOperationNotFound
	}
	return e, nil
}

func (r *operationRepository) List(ctx context.Context, filter repository.OperationFilter) ([]*operationevent.OperationEvent, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*operationevent.OperationEvent
	for _, e := range r.items {
		if e.FarmID != filter.FarmID {
			continue
		}
		if filter.Type != nil && e.Type != *filter.Type {
			continue
		}
		result = append(result, e)
	}
	return result, nil
}

var _ repository.OperationRepository = (*operationRepository)(nil)
