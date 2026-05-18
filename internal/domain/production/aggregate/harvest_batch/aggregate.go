package harvestbatch

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(cycleID vo.ID, unitID vo.ID, q vo.Quantity, grade QualityGrade) *Aggregate {
	now := time.Now()

	root := HarvestBatch{
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

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewHarvestBatchCreated(root.ID))

	return a
}
