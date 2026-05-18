package operationevent

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(farmID vo.ID, opType OperationType) *Aggregate {
	root := OperationEvent{
		ID:        vo.NewID(),
		FarmID:    farmID,
		Type:      opType,
		Timestamp: time.Now(),
		Payload:   make(Payload),
		Metadata:  vo.NewMetadata(),
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewOperationCreated(root.ID))

	return a
}
