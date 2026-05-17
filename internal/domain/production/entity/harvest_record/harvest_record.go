package harvestrecord

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type HarvestRecord struct {
	ID vo.ID

	CycleID vo.ID

	Quantity vo.Quantity

	Grade *string

	Notes *string

	HarvestedAt time.Time
}
