package operations

import (
	"context"
	"sync"

	"github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/task"
	"github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type taskRepository struct {
	mu    sync.RWMutex
	items map[vo.ID]*task.Task
}

func NewTaskRepository() repository.TaskRepository {
	return &taskRepository{items: make(map[vo.ID]*task.Task)}
}

func (r *taskRepository) Save(ctx context.Context, t *task.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[t.ID] = t
	return nil
}

func (r *taskRepository) GetByID(ctx context.Context, id vo.ID) (*task.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.items[id]
	if !ok {
		return nil, task.ErrTaskNotFound
	}
	return t, nil
}

func (r *taskRepository) List(ctx context.Context, filter repository.TaskFilter) ([]*task.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*task.Task
	for _, t := range r.items {
		if t.FarmID != filter.FarmID {
			continue
		}
		if filter.Status != nil && t.Status != *filter.Status {
			continue
		}
		result = append(result, t)
	}
	return result, nil
}

func (r *taskRepository) Delete(ctx context.Context, id vo.ID) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.items, id)
	return nil
}

var _ repository.TaskRepository = (*taskRepository)(nil)
