package disease

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Disease struct {
	ev.BaseAggregate
	ID             vo.ID
	Name           string
	ScientificName *string
	PathogenType   PathogenType
	Hosts          []Host
	Symptoms       []Symptom
	Description    *string
	Metadata       vo.Metadata
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ArchivedAt     *time.Time
}

func New(name string, pathogen PathogenType) *Disease {
	now := time.Now()

	root := &Disease{
		ID:           vo.NewID(),
		Name:         name,
		PathogenType: pathogen,
		Hosts:        make([]Host, 0),
		Symptoms:     make([]Symptom, 0),
		Metadata:     vo.NewMetadata(),
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	root.AddEvent(NewDiseaseCreated(root.ID))

	return root
}

func (a *Disease) AddHost(cropID vo.ID) {
	a.Hosts = append(a.Hosts, Host{CropID: cropID})
	a.UpdatedAt = time.Now()

	a.AddEvent(NewHostAdded(a.ID, cropID))
}

func (a *Disease) AddSymptom(s Symptom) {
	a.Symptoms = append(a.Symptoms, s)
	a.UpdatedAt = time.Now()

	a.AddEvent(NewSymptomAdded(a.ID))
}

func (a *Disease) Archive() {
	now := time.Now()

	a.ArchivedAt = &now
	a.UpdatedAt = now

	a.AddEvent(NewDiseaseArchived(a.ID))
}
