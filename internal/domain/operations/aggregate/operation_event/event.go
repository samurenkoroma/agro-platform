package operationevent

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type OperationEvent struct {
	ev.BaseAggregate
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

func New(farmID vo.ID, opType OperationType) *OperationEvent {
	root := &OperationEvent{
		ID:        vo.NewID(),
		FarmID:    farmID,
		Type:      opType,
		Timestamp: time.Now(),
		Payload:   make(Payload),
		Metadata:  vo.NewMetadata(),
	}

	root.AddEvent(NewOperationRecorded(root.ID))

	return root
}
