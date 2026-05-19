package yieldrecord

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventYieldRecorded = "yield.recorded"
	EventYieldArchived = "yield.archived"
)

type YieldRecorded struct {
	ev.BaseEvent
}

func NewYieldRecorded(id vo.ID) YieldRecorded {
	return YieldRecorded{ev.NewBaseEvent(id, EventYieldRecorded)}
}

type YieldArchived struct {
	ev.BaseEvent
}

func NewYieldArchived(id vo.ID) YieldArchived {
	return YieldArchived{ev.NewBaseEvent(id, EventYieldArchived)}
}
