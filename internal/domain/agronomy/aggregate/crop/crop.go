package crop

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Crop struct {
	ev.BaseAggregate
	ID                vo.ID
	Name              string
	ScientificName    string
	Family            string
	Category          CropCategory
	DefaultProtocolID *vo.ID
	Metadata          vo.Metadata
	CreatedAt         time.Time
	UpdatedAt         time.Time
	ArchivedAt        *time.Time
}

func New(name string, category CropCategory, family string, scientificName string) *Crop {
	now := time.Now()

	root := &Crop{
		ID:             vo.NewID(),
		Family:         family,
		ScientificName: scientificName,
		Name:           name,
		Category:       category,
		Metadata:       vo.NewMetadata(),
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	root.AddEvent(NewCropCreated(root.ID))

	return root
}

func (a *Crop) Rename(name string) {
	a.Name = name
	a.UpdatedAt = time.Now()

	a.AddEvent(NewCropRenamed(a.ID))
}

func (a *Crop) AssignProtocol(id vo.ID) {
	a.DefaultProtocolID = &id
	a.UpdatedAt = time.Now()

	a.AddEvent(NewProtocolAssigned(a.ID, id))
}

func (a *Crop) Archive() {
	now := time.Now()

	a.ArchivedAt = &now
	a.UpdatedAt = now

	a.AddEvent(NewCropArchived(a.ID))
}
