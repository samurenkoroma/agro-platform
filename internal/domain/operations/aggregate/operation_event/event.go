package operationevent

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type OperationEvent struct {
	ID               vo.ID
	Type             OperationType
	FarmID           vo.ID
	ProductionUnitID *vo.ID
	GrowingCycleID   *vo.ID
	PlantID          *vo.ID
	HarvestBatchID   *vo.ID
	YieldBatchID     *vo.ID
	PerformedBy      *vo.ID
	Timestamp        time.Time
	Payload          Payload
	Metadata         vo.Metadata
}

type Aggregate struct {
	ev.AggregateRoot
	Root OperationEvent
}
