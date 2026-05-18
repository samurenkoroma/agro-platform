package yieldbatch

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(cycleID vo.ID, plantID vo.ID, q vo.Quantity, grade QualityGrade) *Aggregate {
	now := time.Now()

	root := YieldBatch{
		ID:             vo.NewID(),
		GrowingCycleID: cycleID,
		PlantID:        plantID,
		Quantity:       q,
		Grade:          grade,
		Marketable:     true,
		HarvestedAt:    now,
		CreatedAt:      now,
		Metadata:       vo.NewMetadata(),
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewYieldBatchCreated(root.ID))

	return a
}
