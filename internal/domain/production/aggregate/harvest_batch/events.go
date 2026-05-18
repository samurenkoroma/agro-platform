package harvestbatch

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventHarvestBatchCreated = "harvest.batch.created"

	EventHarvestBatchRejected = "harvest.batch.rejected"

	EventHarvestBatchSold = "harvest.batch.sold"
)

type HarvestBatchCreated struct {
	ev.BaseEvent
}

func NewHarvestBatchCreated(id vo.ID) HarvestBatchCreated {
	return HarvestBatchCreated{
		BaseEvent: ev.NewBaseEvent(id, EventHarvestBatchCreated),
	}
}
