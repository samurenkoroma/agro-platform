package crop

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Crop struct {
	ID                vo.ID
	Name              string
	ScientificName    *string
	Category          CropCategory
	DefaultProtocolID *vo.ID
	Metadata          vo.Metadata
	CreatedAt         time.Time
	UpdatedAt         time.Time
	ArchivedAt        *time.Time
}

type Aggregate struct {
	ev.AggregateRoot
	Root Crop
}
