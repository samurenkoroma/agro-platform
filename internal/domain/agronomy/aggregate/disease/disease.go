package disease

import (
	"time"

	"github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Disease struct {
	ID         vo.ID
	Name       string
	HostID     vo.ID
	PathogenID *vo.ID
	Severity   Severity
	Symptoms   []Symptom
	Metadata   vo.Metadata
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Aggregate struct {
	aggregate.BaseAggregate
	Root Disease
}
