package yieldbatch

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type YieldBatch struct {
	ev.AggregateRoot
	ID             vo.ID
	GrowingCycleID vo.ID
	PlantID        vo.ID
	Quantity       vo.Quantity
	FruitCount     *int
	Grade          QualityGrade
	Marketable     bool
	Notes          *string
	HarvestedAt    time.Time
	Metadata       vo.Metadata
	CreatedAt      time.Time
}

func New(cycleID vo.ID, plantID vo.ID, q vo.Quantity, grade QualityGrade) *YieldBatch {
	now := time.Now()

	root := &YieldBatch{
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

	root.AddEvent(NewYieldBatchCreated(root.ID))

	return root
}
