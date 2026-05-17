package slot

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventSlotCreated  = "slot.created"
	EventSlotOccupied = "slot.occupied"
	EventSlotReleased = "slot.released"
	EventSlotBlocked  = "slot.blocked"
)

type SlotCreated struct {
	ev.BaseEvent
}

func NewSlotCreated(id vo.ID) SlotCreated {
	return SlotCreated{
		BaseEvent: ev.NewBaseEvent(id, EventSlotCreated),
	}
}

type SlotOccupied struct {
	ev.BaseEvent
	Count int
}

func NewSlotOccupied(id vo.ID, count int) SlotOccupied {
	return SlotOccupied{
		BaseEvent: ev.NewBaseEvent(id, EventSlotOccupied),
		Count:     count,
	}
}

type SlotReleased struct {
	ev.BaseEvent
	Count int
}

func NewSlotReleased(id vo.ID, count int) SlotReleased {
	return SlotReleased{
		BaseEvent: ev.NewBaseEvent(id, EventSlotReleased),
		Count:     count,
	}
}

type SlotBlocked struct {
	ev.BaseEvent
}

func NewSlotBlocked(id vo.ID) SlotBlocked {
	return SlotBlocked{
		BaseEvent: ev.NewBaseEvent(id, EventSlotBlocked),
	}
}
