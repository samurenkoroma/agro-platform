package crop

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventCropCreated      = "crop.created"
	EventCropRenamed      = "crop.renamed"
	EventProtocolAssigned = "crop.protocol.assigned"
	EventCropArchived     = "crop.archived"
)

type CropCreated struct {
	ev.BaseEvent
}

func NewCropCreated(id vo.ID) CropCreated {
	return CropCreated{
		BaseEvent: ev.NewBaseEvent(id, EventCropCreated),
	}
}

type CropRenamed struct {
	ev.BaseEvent
}

func NewCropRenamed(id vo.ID) CropRenamed {
	return CropRenamed{
		BaseEvent: ev.NewBaseEvent(id, EventCropRenamed),
	}
}

type ProtocolAssigned struct {
	ev.BaseEvent
	ProtocolID vo.ID
}

func NewProtocolAssigned(id vo.ID, protocolId vo.ID) ProtocolAssigned {
	return ProtocolAssigned{
		BaseEvent:  ev.NewBaseEvent(id, EventProtocolAssigned),
		ProtocolID: protocolId,
	}
}

type CropArchived struct {
	ev.BaseEvent
}

func NewCropArchived(id vo.ID) CropArchived {
	return CropArchived{
		BaseEvent: ev.NewBaseEvent(id, EventCropArchived),
	}
}
