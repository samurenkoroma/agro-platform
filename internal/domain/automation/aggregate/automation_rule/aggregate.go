package automationrule

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(farmID vo.ID, name string, trigger Trigger, action Action) *Aggregate {
	now := time.Now()

	root := Rule{
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

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewRuleCreated(root.ID))

	return a
}
