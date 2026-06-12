package repository

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/timeline"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type TimeLineRepository interface {
	Save(ctx context.Context, t *timeline.Timeline) error
	GetByID(ctx context.Context, id vo.ID) (*timeline.Timeline, error)
	GetByOwner(ctx context.Context, farmID vo.ID, cycleID *vo.ID) (*timeline.Timeline, error)
}
