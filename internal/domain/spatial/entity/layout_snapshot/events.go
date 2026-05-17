package layoutsnapshot

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventLayoutCreated = "layout.created"
	EventUnitCaptured  = "layout.unit.captured"
)

type LayoutCreated struct {
	ev.BaseEvent
}

func NewLayoutCreated(id vo.ID) LayoutCreated {

	return LayoutCreated{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventLayoutCreated,
		),
	}
}

type UnitCaptured struct {
	ev.BaseEvent

	UnitID vo.ID
}

func NewUnitCaptured(layoutID vo.ID, unitID vo.ID) UnitCaptured {

	return UnitCaptured{
		BaseEvent: ev.NewBaseEvent(
			layoutID,
			EventUnitCaptured,
		),

		UnitID: unitID,
	}
}
