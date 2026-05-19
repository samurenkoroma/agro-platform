package automationrule

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Rule struct {
	ID        vo.ID
	Name      string
	FarmID    vo.ID
	Status    Status
	Trigger   Trigger
	Action    Action
	Metadata  vo.Metadata
	CreatedAt time.Time
	UpdatedAt time.Time

	ArchivedAt *time.Time
}

type Aggregate struct {
	ev.AggregateRoot
	Root Rule
}
