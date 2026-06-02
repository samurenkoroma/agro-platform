package allocation

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type AllocateProductionUnitCommand struct {
	CycleID          vo.ID
	ProductionUnitID vo.ID
	Area             float64
	StartedAt        time.Time
}
type ChangeAllocationCommand struct {
	ID               vo.ID      `json:"id"`
	ProductionUnitID vo.ID      `json:"production_unit_id"`
	Area             float64    `json:"area"`
	StartedAt        time.Time  `json:"started_at"`
	EndedAt          *time.Time `json:"ended_at"`
}

type ReleaseAllocationCommand struct {
	ID         vo.ID      `json:"id"`
	ReleasedAt *time.Time `json:"released_at"`
}
