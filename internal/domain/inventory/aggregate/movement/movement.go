package movement

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Movement struct {
	ID        vo.ID
	ItemID    vo.ID
	Type      Type
	Quantity  float64
	Reference *Reference
	Timestamp time.Time
	Metadata  vo.Metadata
}

type Aggregate struct {
	ev.AggregateRoot

	Root Movement
}
