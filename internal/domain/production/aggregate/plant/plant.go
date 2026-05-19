package plant

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Plant struct {
	ev.AggregateRoot
	ID               vo.ID
	GrowingCycleID   vo.ID
	CropID           vo.ID
	VarietyID        *vo.ID
	ProductionUnitID vo.ID
	SlotID           *vo.ID
	SubstrateID      *vo.ID
	Status           PlantStatus
	Health           PlantHealth
	CurrentStageID   *vo.ID
	PlantedAt        time.Time
	TransplantedAt   *time.Time
	HarvestedAt      *time.Time
	DiscardedAt      *time.Time
	Metadata         vo.Metadata
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func New(cycleID vo.ID, cropID vo.ID, unitID vo.ID) *Plant {

	now := time.Now()

	root := &Plant{
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

	root.AddEvent(
		NewPlantCreated(
			root.ID,
		),
	)

	return root
}

func (a *Plant) Transplant(unitID vo.ID, slotID *vo.ID) error {
	if a.Status == Harvested {

		return ErrAlreadyHarvested
	}

	if a.Status == Dead {

		return ErrPlantDead
	}

	now := time.Now()
	a.ProductionUnitID = unitID
	a.SlotID = slotID
	a.Status = Transplanted
	a.TransplantedAt = &now
	a.UpdatedAt = now

	a.AddEvent(NewPlantTransplanted(a.ID, unitID))

	return nil
}

func (a *Plant) SetStress() {
	a.Health = Stress
	a.Status = Stressed
	a.AddEvent(NewPlantStressed(a.ID))
}

func (a *Plant) SetDisease() {
	a.Health = Disease
	a.Status = Diseased

	a.AddEvent(NewPlantDiseased(a.ID))
}

func (a *Plant) Discard() error {
	if a.Status == Harvested {
		return ErrAlreadyHarvested
	}

	now := time.Now()

	a.Status = Discarded
	a.DiscardedAt = &now
	a.UpdatedAt = now

	a.AddEvent(NewPlantDiscarded(a.ID))

	return nil
}

func (a *Plant) Harvest() error {
	if a.Status == Harvested {

		return ErrAlreadyHarvested
	}

	now := time.Now()
	a.Status = Harvested
	a.HarvestedAt = &now
	a.UpdatedAt = now

	a.AddEvent(NewPlantHarvested(a.ID))

	return nil
}

func (a *Plant) Die() error {
	if a.Status == Harvested {

		return ErrAlreadyHarvested
	}

	now := time.Now()
	a.Status = Dead
	a.UpdatedAt = now

	a.AddEvent(NewPlantDied(a.ID))

	return nil
}
