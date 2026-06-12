package operations

import (
	"context"
	"sync"

	"github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/timeline"
	"github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type timelineRepository struct {
	mu    sync.RWMutex
	items map[vo.ID]*timeline.Timeline
}

func NewTimelineRepository() repository.TimeLineRepository {
	return &timelineRepository{items: make(map[vo.ID]*timeline.Timeline)}
}

func (r *timelineRepository) Save(ctx context.Context, t *timeline.Timeline) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[t.ID] = t
	return nil
}

func (r *timelineRepository) GetByID(ctx context.Context, id vo.ID) (*timeline.Timeline, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.items[id]
	if !ok {
		return nil, nil
	}
	return t, nil
}

func (r *timelineRepository) GetByOwner(ctx context.Context, farmID vo.ID, cycleID *vo.ID) (*timeline.Timeline, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, t := range r.items {
		if t.FarmID != farmID {
			continue
		}
		if cycleID != nil && (t.GrowingCycleID == nil || *t.GrowingCycleID != *cycleID) {
			continue
		}
		return t, nil
	}
	return nil, nil
}

var _ repository.TimeLineRepository = (*timelineRepository)(nil)
