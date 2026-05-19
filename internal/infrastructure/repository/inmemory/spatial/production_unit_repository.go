package spatial

import (
	"sync"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
)

type ProductionUnitRepository struct {
	mu sync.RWMutex

	items map[vo.ID]*pu.ProductionUnit
}

func NewProductionUnitRepository() *ProductionUnitRepository {
	return &ProductionUnitRepository{
		items: make(map[vo.ID]*pu.ProductionUnit),
	}
}
