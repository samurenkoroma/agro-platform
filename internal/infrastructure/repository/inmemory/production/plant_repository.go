package production

import (
	"sync"

	"github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/plant"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type PlantRepository struct {
	mu    sync.RWMutex
	items map[vo.ID]*plant.Plant
}

func NewPlantRepository() *PlantRepository {
	return &PlantRepository{
		items: make(map[vo.ID]*plant.Plant),
	}
}
