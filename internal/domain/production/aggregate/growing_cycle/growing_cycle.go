package growingcycle

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type GrowingCycle struct {
	ev.AggregateRoot
	ID                vo.ID
	FarmID            vo.ID
	CropID            vo.ID
	VarietyID         *vo.ID
	ProductionUnitID  vo.ID
	Method            GrowingMethod
	Granularity       ProductionGranularity
	ProtocolID        *vo.ID
	Status            GrowingStatus
	CurrentStageID    *vo.ID
	LayoutSnapshotID  *vo.ID
	ExpectedHarvestAt *time.Time
	StartedAt         *time.Time
	CompletedAt       *time.Time
	Metadata          vo.Metadata
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func New(farmID vo.ID, cropID vo.ID, unitID vo.ID, method GrowingMethod) *GrowingCycle {
	root := &GrowingCycle{
		ID:               vo.NewID(),
		FarmID:           farmID,
		CropID:           cropID,
		ProductionUnitID: unitID,
		Method:           method,
		Status:           Planned,
		Metadata:         vo.NewMetadata(),
	}

	root.AddEvent(NewCycleCreated(root.ID))

	return root
}

func (a *GrowingCycle) Start() error {

	if a.Status != Planned {

		return ErrAlreadyStarted
	}

	a.Status = Active
	now := time.Now()
	a.StartedAt = &now

	a.AddEvent(NewCycleStarted(a.ID))

	return nil
}

func (a *GrowingCycle) Pause() error {

	if a.Status != Active {

		return ErrInvalidTransition
	}
	a.Status = Paused

	a.AddEvent(NewCyclePaused(a.ID))

	return nil
}

func (a *GrowingCycle) Resume() error {
	if a.Status != Paused {

		return ErrInvalidTransition
	}

	a.Status = Active
	a.AddEvent(NewCycleResumed(a.ID))

	return nil
}
func (a *GrowingCycle) StartHarvest() error {
	if a.Status != Active {
		return ErrInvalidTransition
	}

	a.Status = Harvesting
	a.AddEvent(NewHarvestStarted(a.ID))

	return nil
}

func (a *GrowingCycle) CompleteHarvest() error {
	if a.Status != Harvesting {

		return ErrInvalidTransition
	}

	now := time.Now()
	a.Status = Harvested
	a.CompletedAt = &now

	a.AddEvent(NewHarvestCompleted(a.ID))

	return nil
}

func (a *GrowingCycle) Fail() error {
	if a.Status != Active {

		return ErrInvalidTransition
	}
	now := time.Now()

	a.Status = Failed
	a.CompletedAt = &now

	a.AddEvent(NewCycleFailed(a.ID))

	return nil
}

func (a *GrowingCycle) Archive() error {
	if a.Status != Harvested && a.Status != Failed {

		return ErrInvalidTransition
	}
	a.Status = Archived
	a.AddEvent(NewCycleArchived(a.ID))

	return nil
}
