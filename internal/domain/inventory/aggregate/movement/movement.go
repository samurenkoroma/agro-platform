package movement

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Movement struct {
	ev.BaseAggregate
	ID        vo.ID
	ItemID    vo.ID
	Type      Type
	Quantity  float64
	Reference *Reference
	Timestamp time.Time
	Metadata  vo.Metadata
}

func New(itemID vo.ID, mType Type, quantity float64) *Movement {
	root := &Movement{
		ID:        vo.NewID(),
		ItemID:    itemID,
		Type:      mType,
		Quantity:  quantity,
		Timestamp: time.Now(),
		Metadata:  vo.NewMetadata(),
	}

	root.AddEvent(NewMovementCreated(root.ID))

	return root
}
