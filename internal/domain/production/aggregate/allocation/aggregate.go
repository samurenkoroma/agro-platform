package allocation

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Allocation struct {
	ev.BaseAggregate
	ID               vo.ID
	CycleID          vo.ID
	ProductionUnitID vo.ID
	Area             float64
	StartedAt        *time.Time
	EndedAt          *time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func New(cycleID, productionUnitID vo.ID, area float64, startedAt *time.Time) *Allocation {
	now := time.Now()
	a := &Allocation{
		ID:               vo.NewID(),
		CycleID:          cycleID,
		ProductionUnitID: productionUnitID,
		Area:             area,
		StartedAt:        startedAt,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	a.AddEvent(NewAllocationAllocated(a.ID, productionUnitID))
	return a
}
