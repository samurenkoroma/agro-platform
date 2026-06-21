package variety

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Variety struct {
	ev.BaseAggregate
	ID      vo.ID
	CropID  vo.ID
	Name    string
	Breeder *string

	Profile Profile

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Profile struct {
	Maturity vo.Maturity
	Spacing  PlantSpacing
}

type PlantSpacing struct {
	PlantDistanceCM      *float64
	RowDistanceCM        *float64
	PlantsPerSquareMeter *float64
	RecommendedDensity   *float64
}

func New(cropID vo.ID, name string) *Variety {
	now := time.Now()

	root := &Variety{
		ID:        vo.NewID(),
		CropID:    cropID,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	root.AddEvent(NewVarietyCreated(root.ID))

	return root
}
