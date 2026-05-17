package growingcycle

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventCreated = "growing_cycle.created"

	EventStarted = "growing_cycle.started"

	EventPaused = "growing_cycle.paused"

	EventResumed = "growing_cycle.resumed"

	EventHarvested = "growing_cycle.harvested"

	EventFailed = "growing_cycle.failed"

	EventArchived = "growing_cycle.archived"
)

type CycleCreated struct {
	ev.BaseEvent
}

type CycleStarted struct {
	ev.BaseEvent
}
type CyclePaused struct {
	ev.BaseEvent
}
type CycleResumed struct {
	ev.BaseEvent
}
type CycleHarvested struct {
	ev.BaseEvent
}
type CycleFailed struct {
	ev.BaseEvent
}
type CycleArchived struct {
	ev.BaseEvent
}

func NewCycleCreated(
	id vo.ID,
) CycleCreated {

	return CycleCreated{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventCreated,
		),
	}
}
func NewCycleStarted(id vo.ID) CycleStarted {
	return CycleStarted{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventStarted,
		),
	}
}

func NewCyclePaused(id vo.ID) CyclePaused {
	return CyclePaused{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventPaused,
		),
	}
}
func NewCycleResumed(id vo.ID) CycleResumed {
	return CycleResumed{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventResumed,
		),
	}
}
func NewCycleHarvested(id vo.ID) CycleHarvested {
	return CycleHarvested{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventHarvested,
		),
	}
}
func NewCycleFailed(id vo.ID) CycleFailed {
	return CycleFailed{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventFailed,
		),
	}
}
func NewCycleArchived(id vo.ID) CycleArchived {
	return CycleArchived{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventArchived,
		),
	}
}
