package growingcycle

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(
	farmID vo.ID,

	cropID vo.ID,

	unitID vo.ID,

	method GrowingMethod,
) *Aggregate {

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

func (
	a *Aggregate,
) Start() error {

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

func (
	a *Aggregate,
) Pause() error {

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

func (
	a *Aggregate,
) Resume() error {

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
) Harvest() error {

	if a.Root.Status !=
		Active {

		return ErrInvalidTransition
	}

	now := time.Now()

	a.Root.Status =
		Harvested

	a.Root.CompletedAt =
		&now

	a.AddEvent(
		NewCycleHarvested(
			a.Root.ID,
		),
	)

	return nil
}

func (
	a *Aggregate,
) Fail() error {

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

func (
	a *Aggregate,
) Archive() error {

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
