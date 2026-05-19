package yieldrecord

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(farmID vo.ID, cycleID vo.ID, unitID vo.ID, quantity float64) *Aggregate {
	root := YieldRecord{
		ID:               vo.NewID(),
		FarmID:           farmID,
		GrowingCycleID:   cycleID,
		ProductionUnitID: unitID,
		Quantity:         quantity,
		HarvestedAt:      time.Now(),
		Metadata:         vo.NewMetadata(),
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewYieldRecorded(root.ID))

	return a
}
