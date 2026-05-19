package automationrule

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Rule struct {
	ev.AggregateRoot
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

func New(farmID vo.ID, name string, trigger Trigger, action Action) *Rule {
	now := time.Now()

	root := &Rule{
		ID:       vo.NewID(),
		Name:     name,
		FarmID:   farmID,
		Status:   Enabled,
		Trigger:  trigger,
		Action:   action,
		Metadata: vo.NewMetadata(),

		CreatedAt: now,
		UpdatedAt: now,
	}

	root.AddEvent(NewRuleCreated(root.ID))

	return root
}
