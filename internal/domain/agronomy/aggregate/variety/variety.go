package variety

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Variety struct {
	ev.BaseAggregate
	ID        vo.ID
	CropID    vo.ID
	Name      string
	Breeder   *string
	Maturity  MaturityProfile
	Growth    GrowthProfile
	Spacing   PlantSpacing
	Harvest   HarvestProfile
	Yield     YieldPotential
	Tolerance EnvironmentTolerance

	BaseTemperature float64 // Tbase (ниже которой рост останавливается)
	MaxTemperature  float64 // Tmax (выше которой рост не ускоряется)

	// Фенология (GDD требования)
	PhenophaseGDD []PhenophaseGDD `json:"phenophaseGDD"`
	// Водные требования
	WaterRequirement WaterRequirement `json:"water_requirement"`
	// Световые требования
	LightRequirement LightRequirement `json:"light_requirement"`

	// Характеристики
	Characteristics map[string]string `json:"characteristics"`
	Image           string            `json:"image"`

	Metadata   vo.Metadata
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ArchivedAt *time.Time
}

func New(cropID vo.ID, name string) *Variety {
	now := time.Now()

	root := &Variety{
		ID:               vo.NewID(),
		CropID:           cropID,
		Name:             name,
		Metadata:         vo.NewMetadata(),
		Characteristics:  make(map[string]string),
		PhenophaseGDD:    make([]PhenophaseGDD, 0),
		WaterRequirement: WaterRequirement{},
		LightRequirement: LightRequirement{},

		Image:     "",
		CreatedAt: now,
		UpdatedAt: now,
	}

	root.AddEvent(NewVarietyCreated(root.ID))

	return root
}

func (a *Variety) UpdateMaturity(m MaturityProfile) {
	a.Maturity = m
	a.UpdatedAt = time.Now()

	a.AddEvent(NewMaturityUpdated(a.ID))
}

func (a *Variety) UpdateGrowth(g GrowthProfile) {
	a.Growth = g

	a.AddEvent(NewGrowthUpdated(a.ID))
}

func (a *Variety) UpdateHarvest(h HarvestProfile) {
	a.Harvest = h

	a.AddEvent(NewHarvestUpdated(a.ID))
}

func (a *Variety) UpdateYield(y YieldPotential) {
	a.Yield = y

	a.AddEvent(NewYieldUpdated(a.ID))
}

func (a *Variety) UpdateTolerance(t EnvironmentTolerance) {
	a.Tolerance = t

	a.AddEvent(NewToleranceUpdated(a.ID))
}

func (a *Variety) Archive() {
	now := time.Now()
	a.ArchivedAt = &now
	a.UpdatedAt = now

	a.AddEvent(NewVarietyArchived(a.ID))
}
