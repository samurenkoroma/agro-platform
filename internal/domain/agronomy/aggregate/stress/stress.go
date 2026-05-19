package stress

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Stress struct {
	ID          vo.ID
	Name        string
	Type        StressType
	Description *string
	Triggers    []Trigger
	Symptoms    []Symptom
	Metadata    vo.Metadata
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ArchivedAt  *time.Time
}

type Aggregate struct {
	ev.AggregateRoot
	Root Stress
}
