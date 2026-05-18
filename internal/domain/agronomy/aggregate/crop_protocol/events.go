package cropprotocol

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventProtocolCreated  = "crop.protocol.created"
	EventStageAdded       = "crop.protocol.stage.added"
	EventStageUpdated     = "crop.protocol.stage.updated"
	EventProtocolArchived = "crop.protocol.archived"
)

type ProtocolCreated struct {
	ev.BaseEvent
}

func NewProtocolCreated(id vo.ID) ProtocolCreated {
	return ProtocolCreated{
		BaseEvent: ev.NewBaseEvent(id, EventProtocolCreated),
	}
}

type StageAdded struct {
	ev.BaseEvent
}

func NewStageAdded(id vo.ID) StageAdded {
	return StageAdded{
		BaseEvent: ev.NewBaseEvent(id, EventStageAdded),
	}
}

type StageUpdated struct {
	ev.BaseEvent
}

func NewStageUpdated(id vo.ID) StageUpdated {
	return StageUpdated{
		BaseEvent: ev.NewBaseEvent(id, EventStageUpdated),
	}
}

type ProtocolArchived struct {
	ev.BaseEvent
}

func NewProtocolArchived(id vo.ID) ProtocolArchived {
	return ProtocolArchived{
		BaseEvent: ev.NewBaseEvent(id, EventProtocolArchived),
	}
}
