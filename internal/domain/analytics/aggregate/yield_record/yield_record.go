package yieldrecord

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type YieldRecord struct {
	ev.AggregateRoot
	ID               vo.ID
	FarmID           vo.ID
	GrowingCycleID   vo.ID
	ProductionUnitID vo.ID
	HarvestBatchID   *vo.ID
	YieldBatchID     *vo.ID
	Source           Source
	Quantity         float64
	Unit             string
	HarvestedAt      time.Time
	Metadata         vo.Metadata
}

func New(farmID vo.ID, cycleID vo.ID, unitID vo.ID, quantity float64) *YieldRecord {
	root := &YieldRecord{
		ID:               vo.NewID(),
		FarmID:           farmID,
		GrowingCycleID:   cycleID,
		ProductionUnitID: unitID,
		Quantity:         quantity,
		HarvestedAt:      time.Now(),
		Metadata:         vo.NewMetadata(),
	}

	root.AddEvent(NewYieldRecorded(root.ID))

	return root
}
