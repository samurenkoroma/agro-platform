package crop

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(name string, category CropCategory) *Aggregate {
	now := time.Now()

	root := Crop{
		ID:        vo.NewID(),
		Name:      name,
		Category:  category,
		Metadata:  vo.NewMetadata(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewCropCreated(root.ID))

	return a
}

func (a *Aggregate) Rename(name string) {
	a.Root.Name = name

	a.Root.UpdatedAt = time.Now()

	a.AddEvent(NewCropRenamed(a.Root.ID))
}

func (a *Aggregate) AssignProtocol(id vo.ID) {

	a.Root.DefaultProtocolID = &id

	a.Root.UpdatedAt = time.Now()

	a.AddEvent(NewProtocolAssigned(a.Root.ID, id))
}

func (a *Aggregate) Archive() {
	now := time.Now()

	a.Root.ArchivedAt = &now

	a.Root.UpdatedAt = now

	a.AddEvent(NewCropArchived(a.Root.ID))
}
