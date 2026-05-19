package disease

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Disease struct {
	ID             vo.ID
	Name           string
	ScientificName *string
	PathogenType   PathogenType
	Hosts          []Host
	Symptoms       []Symptom
	Description    *string
	Metadata       vo.Metadata
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ArchivedAt     *time.Time
}

type Aggregate struct {
	ev.AggregateRoot
	Root Disease
}
