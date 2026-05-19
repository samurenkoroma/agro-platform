package climatezone

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type ClimateZone struct {
	ev.AggregateRoot
	ID         vo.ID
	FarmID     vo.ID
	Name       string
	Target     Target
	Metadata   vo.Metadata
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ArchivedAt *time.Time
}

func New(farmID vo.ID, name string) *ClimateZone {
	now := time.Now()

	root := &ClimateZone{
		ID:       vo.NewID(),
		FarmID:   farmID,
		Name:     name,
		Metadata: vo.NewMetadata(),

		CreatedAt: now,
		UpdatedAt: now,
	}

	root.AddEvent(NewClimateZoneCreated(root.ID))

	return root
}

func (a *ClimateZone) UpdateTarget(target Target) error {

	if a.ArchivedAt != nil {
		return ErrArchivedZone
	}

	a.Target = target
	a.UpdatedAt = time.Now()

	a.AddEvent(NewClimateZoneTargetUpdated(a.ID))

	return nil
}
