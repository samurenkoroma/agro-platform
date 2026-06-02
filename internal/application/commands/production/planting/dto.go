package planting

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type RegisterPlantingCommand struct {
	CycleID   vo.ID     `json:"cycle_id"`
	PlantedAt time.Time `json:"planted_at"`
	Quantity  float64   `json:"quantity"`
}

type ChangePlantingCommand struct {
	ID        vo.ID     `json:"id"`
	PlantedAt time.Time `json:"planted_at"`
	Quantity  float64   `json:"quantity"`
}
