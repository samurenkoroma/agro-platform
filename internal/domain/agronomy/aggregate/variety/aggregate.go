package variety

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(cropID vo.ID, name string) *Aggregate {
	now := time.Now()

	root := Variety{
		ID:        vo.NewID(),
		CropID:    cropID,
		Name:      name,
		Metadata:  vo.NewMetadata(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewVarietyCreated(root.ID))

	return a
}

func (a *Aggregate) UpdateMaturity(m MaturityProfile) {
	a.Root.Maturity = m

	a.Root.UpdatedAt = time.Now()

	a.AddEvent(NewMaturityUpdated(a.Root.ID))
}

func (a *Aggregate) UpdateGrowth(g GrowthProfile) {
	a.Root.Growth = g

	a.AddEvent(NewGrowthUpdated(a.Root.ID))
}
func (a *Aggregate) UpdateHarvest(h HarvestProfile) {
	a.Root.Harvest = h

	a.AddEvent(NewHarvestUpdated(a.Root.ID))
}
func (a *Aggregate) UpdateYield(y YieldPotential) {
	a.Root.Yield = y

	a.AddEvent(NewYieldUpdated(a.Root.ID))
}
func (a *Aggregate) UpdateTolerance(t EnvironmentTolerance) {
	a.Root.Tolerance = t

	a.AddEvent(NewToleranceUpdated(a.Root.ID))
}
func (a *Aggregate) Archive() {
	now := time.Now()

	a.Root.ArchivedAt = &now
	a.Root.UpdatedAt = now

	a.AddEvent(NewVarietyArchived(a.Root.ID))
}
