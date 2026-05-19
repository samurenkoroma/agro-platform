package stress

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Stress struct {
	ev.AggregateRoot
	ID          vo.ID
	Name        string
	Type        StressType
	Description *string
	Triggers    []Trigger
	Symptoms    []Symptom
	Metadata    vo.Metadata
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ArchivedAt  *time.Time
}

func New(name string, stressType StressType) *Stress {
	now := time.Now()

	root := &Stress{
		ID:        vo.NewID(),
		Name:      name,
		Type:      stressType,
		Triggers:  make([]Trigger, 0),
		Symptoms:  make([]Symptom, 0),
		Metadata:  vo.NewMetadata(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	root.AddEvent(NewStressCreated(root.ID))

	return root
}

func (a *Stress) AddTrigger(trigger Trigger) {
	a.Triggers = append(a.Triggers, trigger)
	a.UpdatedAt = time.Now()

	a.AddEvent(NewTriggerAdded(a.ID))
}

func (a *Stress) AddSymptom(s Symptom) {
	a.Symptoms = append(a.Symptoms, s)
	a.UpdatedAt = time.Now()

	a.AddEvent(NewSymptomAdded(a.ID))
}

func (a *Stress) Archive() {
	now := time.Now()
	a.ArchivedAt = &now
	a.UpdatedAt = now

	a.AddEvent(NewStressArchived(a.ID))
}
