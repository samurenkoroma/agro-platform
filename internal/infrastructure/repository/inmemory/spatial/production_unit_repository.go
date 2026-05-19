package spatial

import (
	"sync"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
	"github.com/samurenkoroma/agro-platform/internal/domain/spatial/repository"
)

type productionUnitRepository struct {
	mu sync.RWMutex

	items map[vo.ID]*pu.ProductionUnit
}

func (r *productionUnitRepository) Save(aggregate *pu.ProductionUnit) error {
	//TODO implement me
	panic("implement me")
}

func (r *productionUnitRepository) GetByID(id vo.ID) (*pu.ProductionUnit, error) {
	//TODO implement me
	panic("implement me")
}

func (r *productionUnitRepository) GetChildren(parentID vo.ID) ([]*pu.ProductionUnit, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*pu.ProductionUnit, 0)

	for _, unit := range r.items {
		if unit.ParentID == nil {
			continue
		}

		if *unit.ParentID == parentID {
			result = append(result, unit)
		}
	}

	return result,
		nil
}

func (r *productionUnitRepository) Delete(id vo.ID) error {
	//TODO implement me
	panic("implement me")
}

func NewProductionUnitRepository() repository.ProductionUnitRepository {
	return &productionUnitRepository{
		items: make(map[vo.ID]*pu.ProductionUnit),
	}
}
