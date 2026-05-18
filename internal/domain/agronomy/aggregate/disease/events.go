package disease

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventDiseaseCreated  = "disease.created"
	EventHostAdded       = "disease.host.added"
	EventSymptomAdded    = "disease.symptom.added"
	EventDiseaseArchived = "disease.archived"
)

type DiseaseCreated struct {
	ev.BaseEvent
}

func NewDiseaseCreated(id vo.ID) DiseaseCreated {
	return DiseaseCreated{
		BaseEvent: ev.NewBaseEvent(id, EventDiseaseCreated),
	}
}

type HostAdded struct {
	ev.BaseEvent
	CropID vo.ID
}

func NewHostAdded(id vo.ID, cropID vo.ID) HostAdded {
	return HostAdded{
		BaseEvent: ev.NewBaseEvent(id, EventHostAdded),
		CropID:    cropID,
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

type DiseaseArchived struct {
	ev.BaseEvent
}

func NewDiseaseArchived(id vo.ID) DiseaseArchived {
	return DiseaseArchived{
		BaseEvent: ev.NewBaseEvent(id, EventDiseaseArchived),
	}
}
