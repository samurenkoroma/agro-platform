package allocation

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type AllocateProductionUnitCommand struct {
	CycleID          vo.ID      `json:"cycleID" validate:"required"`
	ProductionUnitID vo.ID      `json:"productionUnitID" validate:"required"`
	Area             float64    `json:"area" validate:"required"`
	StartedAt        *time.Time `json:"startedAt,omitempty"`
}
type ChangeAllocationCommand struct {
	ID               vo.ID      `json:"id"`
	ProductionUnitID vo.ID      `json:"productionUnitID"`
	Area             float64    `json:"area"`
	StartedAt        *time.Time `json:"startedAt"`
	EndedAt          *time.Time `json:"endedAt"`
}

type ReleaseAllocationCommand struct {
	ID         vo.ID      `json:"id"`
	ReleasedAt *time.Time `json:"released_at"`
}
