package operationevent

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const EventOperationRecorded = "operation.recorded"

type OperationRecorded struct{ ev.BaseEvent }

func NewOperationRecorded(id vo.ID) OperationRecorded {
	return OperationRecorded{ev.NewBaseEvent(id, EventOperationRecorded)}
}
