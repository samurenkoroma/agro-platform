package growingcycle

import (
	"time"

	harvestrecord "github.com/samurenkoroma/agro-platform/internal/domain/production/entity/harvest_record"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(farmID vo.ID, cropID vo.ID, unitID vo.ID, method GrowingMethod) *Aggregate {

	root := GrowingCycle{
		ID: vo.NewID(),

		FarmID: farmID,

		CropID: cropID,

		ProductionUnitID: unitID,

		Method: method,

		Status: Planned,

		Metadata: vo.NewMetadata(),
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(
		NewCycleCreated(
			root.ID,
		),
	)

	return a
}

func (a *Aggregate) Start() error {

	if a.Root.Status !=
		Planned {

		return ErrAlreadyStarted
	}

	a.Root.Status =
		Active

	a.Root.StartedAt =
		time.Now()

	a.AddEvent(
		NewCycleStarted(
			a.Root.ID,
		),
	)

	return nil
}

func (a *Aggregate) Pause() error {

	if a.Root.Status !=
		Active {

		return ErrInvalidTransition
	}

	a.Root.Status =
		Paused

	a.AddEvent(
		NewCyclePaused(
			a.Root.ID,
		),
	)

	return nil
}

func (a *Aggregate) Resume() error {

	if a.Root.Status !=
		Paused {

		return ErrInvalidTransition
	}

	a.Root.Status =
		Active

	a.AddEvent(
		NewCycleResumed(
			a.Root.ID,
		),
	)

	return nil
}
func (
	a *Aggregate,
) StartHarvest() error {

	if a.Root.Status != Active {

		return ErrInvalidTransition
	}

	a.Root.Status =
		Harvesting

	a.AddEvent(
		NewHarvestStarted(
			a.Root.ID,
		),
	)

	return nil
}

func (
	a *Aggregate,
) AddHarvest(
	q vo.Quantity,

	grade *string,
) error {

	if a.Root.Status != Harvesting &&
		a.Root.Status != Active {

		return ErrInvalidTransition
	}

	record :=
		harvestrecord.HarvestRecord{

			ID: vo.NewID(),

			CycleID: a.Root.ID,

			Quantity: q,

			Grade: grade,

			HarvestedAt: time.Now(),
		}

	a.Root.Harvests =
		append(
			a.Root.Harvests,
			record,
		)

	a.Root.Status =
		Harvesting

	a.AddEvent(
		NewPartialHarvest(a.Root.ID, record.ID),
	)

	return nil
}

func (
	a *Aggregate,
) CompleteHarvest() error {

	if a.Root.Status !=
		Harvesting {

		return ErrInvalidTransition
	}

	now := time.Now()

	a.Root.Status =
		Harvested

	a.Root.CompletedAt =
		&now

	a.AddEvent(
		NewHarvestCompleted(
			a.Root.ID,
		),
	)

	return nil
}

func (a *Aggregate) Fail() error {

	if a.Root.Status !=
		Active {

		return ErrInvalidTransition
	}

	now := time.Now()

	a.Root.Status =
		Failed

	a.Root.CompletedAt =
		&now

	a.AddEvent(
		NewCycleFailed(
			a.Root.ID,
		),
	)

	return nil
}

func (a *Aggregate) Archive() error {

	if a.Root.Status !=
		Harvested &&
		a.Root.Status !=
			Failed {

		return ErrInvalidTransition
	}

	a.Root.Status =
		Archived

	a.AddEvent(
		NewCycleArchived(
			a.Root.ID,
		),
	)

	return nil
}
