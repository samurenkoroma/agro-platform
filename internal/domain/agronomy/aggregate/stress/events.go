package stress

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventStressCreated  = "stress.created"
	EventTriggerAdded   = "stress.trigger.added"
	EventSymptomAdded   = "stress.symptom.added"
	EventStressArchived = "stress.archived"
)

type StressCreated struct {
	ev.BaseEvent
}

func NewStressCreated(id vo.ID) StressCreated {
	return StressCreated{
		BaseEvent: ev.NewBaseEvent(id, EventStressCreated),
	}
}

type StressArchived struct {
	ev.BaseEvent
}

func NewStressArchived(id vo.ID) StressArchived {
	return StressArchived{
		BaseEvent: ev.NewBaseEvent(id, EventStressArchived),
	}
}

type SymptomAdded struct {
	ev.BaseEvent
}

func NewSymptomAdded(id vo.ID) SymptomAdded {
	return SymptomAdded{
		BaseEvent: ev.NewBaseEvent(id, EventSymptomAdded),
	}
}

type TriggerAdded struct {
	ev.BaseEvent
}

func NewTriggerAdded(id vo.ID) TriggerAdded {
	return TriggerAdded{
		BaseEvent: ev.NewBaseEvent(id, EventTriggerAdded),
	}
}
