package productionmetrics

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Metrics struct {
	ID vo.ID

	FarmID vo.ID

	GrowingCycleID *vo.ID

	ProductionUnitID *vo.ID

	Consumption Consumption

	Efficiency Efficiency

	Metadata vo.Metadata

	CalculatedAt time.Time
}

type Aggregate struct {
	ev.AggregateRoot

	Root Metrics
}
