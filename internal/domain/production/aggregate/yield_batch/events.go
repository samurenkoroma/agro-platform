package yieldbatch

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventYieldBatchCreated  = "yield.batch.created"
	EventYieldBatchRejected = "yield.batch.rejected"
	EventYieldBatchSold     = "yield.batch.sold"
)

type YieldBatchCreated struct {
	ev.BaseEvent
}

func NewYieldBatchCreated(id vo.ID) YieldBatchCreated {
	return YieldBatchCreated{
		BaseEvent: ev.NewBaseEvent(id, EventYieldBatchCreated),
	}
}
