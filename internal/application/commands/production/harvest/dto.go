package harvest

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type RegisterHarvestCommand struct {
	CycleID   vo.ID     `json:"cycleID"`
	HarvestAt time.Time `json:"harvestAt"`
	Quantity  float64   `json:"quantity"`
}

type ChangeHarvestCommand struct {
	ID        vo.ID     `json:"id"`
	HarvestAt time.Time `json:"harvestAt"`
	Quantity  float64   `json:"quantity"`
}
