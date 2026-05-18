package stress

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(name string, stressType StressType) *Aggregate {
	now := time.Now()

	root := Stress{
		ID:        vo.NewID(),
		Name:      name,
		Type:      stressType,
		Triggers:  make([]Trigger, 0),
		Symptoms:  make([]Symptom, 0),
		Metadata:  vo.NewMetadata(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewStressCreated(root.ID))

	return a
}

func (a *Aggregate) AddTrigger(trigger Trigger) {

	a.Root.Triggers = append(a.Root.Triggers, trigger)

	a.Root.UpdatedAt = time.Now()

	a.AddEvent(NewTriggerAdded(a.Root.ID))
}

func (a *Aggregate) AddSymptom(s Symptom) {

	a.Root.Symptoms = append(a.Root.Symptoms, s)

	a.Root.UpdatedAt = time.Now()

	a.AddEvent(NewSymptomAdded(a.Root.ID))
}

func (a *Aggregate) Archive() {
	now := time.Now()

	a.Root.ArchivedAt = &now

	a.Root.UpdatedAt = now

	a.AddEvent(NewStressArchived(a.Root.ID))
}
