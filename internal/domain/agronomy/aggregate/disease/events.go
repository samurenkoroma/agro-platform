package disease

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventDiseaseCreated  = "disease.created"
	EventDiseaseObserved = "disease.observed"
	EventDiseaseResolved = "disease.resolved"
)

type DiseaseCreated struct {
	ev.BaseEvent
}

func NewDiseaseCreated(id vo.ID) DiseaseCreated {
	return DiseaseCreated{
		BaseEvent: ev.NewBaseEvent(id, EventDiseaseCreated),
	}
}

type DiseaseObserved struct {
	ev.BaseEvent
	Symptom Symptom
}

func NewDiseaseObserved(id vo.ID, s Symptom) DiseaseObserved {
	return DiseaseObserved{
		BaseEvent: ev.NewBaseEvent(id, EventDiseaseObserved),
		Symptom:   s,
	}
}

type DiseaseResolved struct {
	ev.BaseEvent
}

func NewDiseaseResolved(id vo.ID) DiseaseResolved {
	return DiseaseResolved{
		BaseEvent: ev.NewBaseEvent(id, EventDiseaseResolved),
	}
}
