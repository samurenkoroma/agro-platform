package substrate

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventSubstrateCreated   = "substrate.created"
	EventSubstrateUsed      = "substrate.used"
	EventSubstrateExhausted = "substrate.exhausted"
	EventSubstrateRecycled  = "substrate.recycled"
	EventSubstrateDisposed  = "substrate.disposed"
)

type SubstrateCreated struct {
	ev.BaseEvent
}
type SubstrateUsed struct {
	ev.BaseEvent
}
type SubstrateExhausted struct {
	ev.BaseEvent
}
type SubstrateRecycled struct {
	ev.BaseEvent
}
type SubstrateDisposed struct {
	ev.BaseEvent
}

func NewSubstrateCreated(id vo.ID) SubstrateCreated {
	return SubstrateCreated{
		BaseEvent: ev.NewBaseEvent(id, EventSubstrateCreated),
	}
}
func NewSubstrateUsed(id vo.ID) SubstrateUsed {
	return SubstrateUsed{
		BaseEvent: ev.NewBaseEvent(id, EventSubstrateUsed),
	}
}
func NewSubstrateExhausted(id vo.ID) SubstrateExhausted {
	return SubstrateExhausted{
		BaseEvent: ev.NewBaseEvent(id, EventSubstrateExhausted),
	}
}
func NewSubstrateRecycled(id vo.ID) SubstrateRecycled {
	return SubstrateRecycled{
		BaseEvent: ev.NewBaseEvent(id, EventSubstrateRecycled),
	}
}
func NewSubstrateDisposed(id vo.ID) SubstrateDisposed {
	return SubstrateDisposed{
		BaseEvent: ev.NewBaseEvent(id, EventSubstrateDisposed),
	}
}
