package movement

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(itemID vo.ID, mType Type, quantity float64) *Aggregate {
	root := Movement{
		ID:        vo.NewID(),
		ItemID:    itemID,
		Type:      mType,
		Quantity:  quantity,
		Timestamp: time.Now(),
		Metadata:  vo.NewMetadata(),
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewMovementCreated(root.ID))

	return a
}
