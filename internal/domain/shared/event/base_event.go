package event

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type BaseEvent struct {
	ID        vo.ID
	Aggregate vo.ID
	Type      string
	Occurred  time.Time
}

func NewBaseEvent(aggregate vo.ID, eventType string) BaseEvent {
	return BaseEvent{
		ID:        vo.NewID(),
		Aggregate: aggregate,
		Type:      eventType,
		Occurred:  time.Now(),
	}
}

func (e BaseEvent) EventID() vo.ID {
	return e.ID
}

func (e BaseEvent) AggregateID() vo.ID {
	return e.Aggregate
}

func (e BaseEvent) EventType() string {
	return e.Type
}

func (e BaseEvent) OccurredAt() time.Time {
	return e.Occurred
}
