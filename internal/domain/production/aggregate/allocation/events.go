package allocation

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type AllocationAllocated struct {
	AllocationID vo.ID

	CycleID   vo.ID
	CycleName string

	ProductionUnitID   vo.ID
	ProductionUnitName string

	Area float64

	StartedAt time.Time

	OccurredAt time.Time
}

type AllocationReleased struct {
	AllocationID vo.ID

	CycleID   vo.ID
	CycleName string

	ProductionUnitID   vo.ID
	ProductionUnitName string

	Area float64

	ReleasedAt time.Time

	OccurredAt time.Time
}
