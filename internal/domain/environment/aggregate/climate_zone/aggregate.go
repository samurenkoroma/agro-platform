package climatezone

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(farmID vo.ID, name string) *Aggregate {
	now := time.Now()

	root := ClimateZone{
		ID:       vo.NewID(),
		FarmID:   farmID,
		Name:     name,
		Metadata: vo.NewMetadata(),

		CreatedAt: now,
		UpdatedAt: now,
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewClimateZoneCreated(root.ID))

	return a
}

func (a *Aggregate) UpdateTarget(target Target) error {

	if a.Root.ArchivedAt != nil {
		return ErrArchivedZone
	}

	a.Root.Target = target

	a.Root.UpdatedAt = time.Now()

	a.AddEvent(NewClimateZoneTargetUpdated(a.Root.ID))

	return nil
}
