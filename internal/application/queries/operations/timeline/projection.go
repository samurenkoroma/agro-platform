package timeline

import (
	"context"
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type TimelineItemDTO struct {
	ID          vo.ID     `json:"id"`
	Source      string    `json:"source"`
	ReferenceID vo.ID     `json:"referenceId"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	Timestamp   time.Time `json:"timestamp"`
}

type TimelineDTO struct {
	ID    vo.ID             `json:"id"`
	Items []TimelineItemDTO `json:"items"`
}

type Projection interface {
	Get(ctx context.Context, farmID vo.ID, cycleID *vo.ID) (*TimelineDTO, error)
}
