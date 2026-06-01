package allocation

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Allocation struct {
	ID               vo.ID
	CycleID          vo.ID
	ProductionUnitID vo.ID
	AreaM2           float64
	StartedAt        time.Time
	EndedAt          *time.Time
	CreatedAt        time.Time
}
