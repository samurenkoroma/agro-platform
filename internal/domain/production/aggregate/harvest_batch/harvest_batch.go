package harvestbatch

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type HarvestBatch struct {
	ev.AggregateRoot
	ID               vo.ID
	GrowingCycleID   vo.ID
	ProductionUnitID vo.ID
	Quantity         vo.Quantity
	HarvestedArea    *vo.Area
	Grade            QualityGrade
	Marketable       bool
	Notes            *string
	HarvestedAt      time.Time
	Metadata         vo.Metadata
	CreatedAt        time.Time
}

func New(cycleID vo.ID, unitID vo.ID, q vo.Quantity, grade QualityGrade) *HarvestBatch {
	now := time.Now()

	root := &HarvestBatch{
		ID:               vo.NewID(),
		GrowingCycleID:   cycleID,
		ProductionUnitID: unitID,
		Quantity:         q,
		Grade:            grade,
		Marketable:       true,
		HarvestedAt:      now,
		CreatedAt:        now,
		Metadata:         vo.NewMetadata(),
	}

	root.AddEvent(NewHarvestBatchCreated(root.ID))

	return root
}
