package productionmetrics

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Metrics struct {
	ev.AggregateRoot
	ID               vo.ID
	FarmID           vo.ID
	GrowingCycleID   *vo.ID
	ProductionUnitID *vo.ID
	Consumption      Consumption
	Efficiency       Efficiency
	Metadata         vo.Metadata
	CalculatedAt     time.Time
}

func New(farmID vo.ID) *Metrics {

	root := &Metrics{
		ID: vo.NewID(),

		FarmID: farmID,

		Consumption: Consumption{},

		Efficiency: Efficiency{},

		CalculatedAt: time.Now(),

		Metadata: vo.NewMetadata(),
	}

	root.AddEvent(
		NewMetricsCalculated(
			root.ID,
		),
	)

	return root
}
