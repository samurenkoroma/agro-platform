package productionmetrics

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(
	farmID vo.ID,
) *Aggregate {

	root := Metrics{
		ID: vo.NewID(),

		FarmID: farmID,

		Consumption: Consumption{},

		Efficiency: Efficiency{},

		CalculatedAt: time.Now(),

		Metadata: vo.NewMetadata(),
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(
		NewMetricsCalculated(
			root.ID,
		),
	)

	return a
}
