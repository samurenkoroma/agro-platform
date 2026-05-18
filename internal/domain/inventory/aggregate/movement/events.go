package movement

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventMovementCreated   = "inventory.movement.created"
	EventReferenceAttached = "inventory.reference.attached"
)

type MovementCreated struct {
	ev.BaseEvent
}

func NewMovementCreated(id vo.ID) MovementCreated {
	return MovementCreated{ev.NewBaseEvent(id, EventMovementCreated)}
}

type ReferenceAttached struct {
	ev.BaseEvent
}

func NewReferenceAttached(id vo.ID) ReferenceAttached {
	return ReferenceAttached{ev.NewBaseEvent(id, EventReferenceAttached)}
}
