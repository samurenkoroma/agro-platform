package planting

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Planting struct {
	ID vo.ID

	CycleID vo.ID

	Quantity int

	SourceType SourceType

	PlantedAt time.Time

	Notes string

	CreatedAt time.Time
}
