package disease

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(name string, pathogen PathogenType) *Aggregate {
	now := time.Now()

	root := Disease{
		ID:           vo.NewID(),
		Name:         name,
		PathogenType: pathogen,
		Hosts:        make([]Host, 0),
		Symptoms:     make([]Symptom, 0),
		Metadata:     vo.NewMetadata(),
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewDiseaseCreated(root.ID))

	return a
}

func (a *Aggregate) AddHost(cropID vo.ID) {
	a.Root.Hosts = append(a.Root.Hosts, Host{CropID: cropID})

	a.Root.UpdatedAt = time.Now()

	a.AddEvent(NewHostAdded(a.Root.ID, cropID))
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

	a.AddEvent(NewDiseaseArchived(a.Root.ID))
}
