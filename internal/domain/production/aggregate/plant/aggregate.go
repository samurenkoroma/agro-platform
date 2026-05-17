package plant

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(cycleID vo.ID, cropID vo.ID, unitID vo.ID) *Aggregate {

	now := time.Now()

	root := Plant{
		ID:               vo.NewID(),
		GrowingCycleID:   cycleID,
		CropID:           cropID,
		ProductionUnitID: unitID,
		Status:           Germinating,
		Health:           Healthy,
		Metadata:         vo.NewMetadata(),
		PlantedAt:        now,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(
		NewPlantCreated(
			root.ID,
		),
	)

	return a
}

func (a *Aggregate) Transplant(unitID vo.ID, slotID *vo.ID) error {

	if a.Root.Status ==
		Harvested {

		return ErrAlreadyHarvested
	}

	if a.Root.Status ==
		Dead {

		return ErrPlantDead
	}

	now := time.Now()

	a.Root.ProductionUnitID =
		unitID

	a.Root.SlotID =
		slotID

	a.Root.Status =
		Transplanted

	a.Root.TransplantedAt =
		&now

	a.Root.UpdatedAt =
		now

	a.AddEvent(
		NewPlantTransplanted(
			a.Root.ID,
			unitID,
		),
	)

	return nil
}

func (
	a *Aggregate,
) SetStress() {

	a.Root.Health =
		Stress

	a.Root.Status =
		Stressed

	a.AddEvent(
		NewPlantStressed(
			a.Root.ID,
		),
	)
}

func (
	a *Aggregate,
) SetDisease() {

	a.Root.Health =
		Disease

	a.Root.Status =
		Diseased

	a.AddEvent(
		NewPlantDiseased(
			a.Root.ID,
		),
	)
}

func (
	a *Aggregate,
) Discard() error {

	if a.Root.Status ==
		Harvested {

		return ErrAlreadyHarvested
	}

	now := time.Now()

	a.Root.Status =
		Discarded

	a.Root.DiscardedAt =
		&now

	a.Root.UpdatedAt =
		now

	a.AddEvent(
		NewPlantDiscarded(
			a.Root.ID,
		),
	)

	return nil
}

func (
	a *Aggregate,
) Harvest() error {

	if a.Root.Status ==
		Harvested {

		return ErrAlreadyHarvested
	}

	now := time.Now()

	a.Root.Status =
		Harvested

	a.Root.HarvestedAt =
		&now

	a.Root.UpdatedAt =
		now

	a.AddEvent(
		NewPlantHarvested(
			a.Root.ID,
		),
	)

	return nil
}

func (a *Aggregate) Die() error {

	if a.Root.Status ==
		Harvested {

		return ErrAlreadyHarvested
	}

	now := time.Now()

	a.Root.Status =
		Dead

	a.Root.UpdatedAt =
		now

	a.AddEvent(
		NewPlantDied(
			a.Root.ID,
		),
	)

	return nil
}
