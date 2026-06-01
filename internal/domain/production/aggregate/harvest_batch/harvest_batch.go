package harvestbatch

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type HarvestBatch struct {
	ID           vo.ID
	CycleID      vo.ID
	HarvestDate  time.Time
	WeightKg     float64
	Units        int
	QualityGrade string
	Notes        string
	CreatedAt    time.Time
}
