package climatezone

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type ClimateZone struct {
	ID         vo.ID
	FarmID     vo.ID
	Name       string
	Target     Target
	Metadata   vo.Metadata
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ArchivedAt *time.Time
}

type Aggregate struct {
	ev.AggregateRoot

	Root ClimateZone
}
