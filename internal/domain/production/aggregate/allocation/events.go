package allocation

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventAllocated = "production.allocation.allocated"
	EventReleased  = "production.allocation.released"
)

type AllocationAllocated struct {
	ev.BaseEvent

	ProductionUnitID vo.ID
}

func NewAllocationAllocated(allocationID vo.ID, prodUnitID vo.ID) AllocationAllocated {
	return AllocationAllocated{BaseEvent: ev.NewBaseEvent(allocationID, EventAllocated), ProductionUnitID: prodUnitID}
}

type AllocationReleased struct {
	ev.BaseEvent

	AllocationID vo.ID

	CycleID   vo.ID
	CycleName string

	ProductionUnitID   vo.ID
	ProductionUnitName string

	Area float64

	ReleasedAt time.Time
}

func NewAllocationReleased(
	allocationID, cycleID vo.ID, cycleName string,
	unitID vo.ID, unitName string,
	area float64, releasedAt time.Time,
) AllocationReleased {
	return AllocationReleased{
		BaseEvent:          ev.NewBaseEvent(allocationID, EventReleased),
		AllocationID:       allocationID,
		CycleID:            cycleID,
		CycleName:          cycleName,
		ProductionUnitID:   unitID,
		ProductionUnitName: unitName,
		Area:               area,
		ReleasedAt:         releasedAt,
	}
}
