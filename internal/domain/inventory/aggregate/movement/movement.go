package movement

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Movement struct {
	ev.BaseAggregate
	ID              vo.ID
	FarmID          vo.ID
	ItemID          vo.ID
	Type            Type
	Quantity        float64
	FromWarehouseID *vo.ID
	ToWarehouseID   *vo.ID
	Reference       *Reference
	Note            *string
	Timestamp       time.Time
	Metadata        vo.Metadata
}

func New(farmID, itemID vo.ID, mType Type, quantity float64) *Movement {
	root := &Movement{
		ID:        vo.NewID(),
		FarmID:    farmID,
		ItemID:    itemID,
		Type:      mType,
		Quantity:  quantity,
		Timestamp: time.Now(),
		Metadata:  vo.NewMetadata(),
	}
	root.AddEvent(NewMovementCreated(root.ID))
	return root
}

func (m *Movement) AttachReference(ref Reference) {
	m.Reference = &ref
	m.AddEvent(NewReferenceAttached(m.ID))
}
