package substrate

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Substrate struct {
	ID             vo.ID
	Name           string
	Type           SubstrateType
	Reusable       bool
	Status         SubstrateStatus
	VolumeLiters   *float64
	WaterRetention *float64
	Aeration       *float64
	Manufacturer   *string
	BatchID        *vo.ID
	Metadata       vo.Metadata
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Aggregate struct {
	ev.AggregateRoot
	Root Substrate
}
