package event

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type DomainEvent interface {
	EventID() vo.ID

	AggregateID() vo.ID

	EventType() string

	OccurredAt() time.Time
}
