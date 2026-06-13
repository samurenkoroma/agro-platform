package inventory

import (
	"context"
	"sync"

	"github.com/samurenkoroma/agro-platform/internal/domain/inventory/aggregate/warehouse"
	"github.com/samurenkoroma/agro-platform/internal/domain/inventory/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type warehouseRepository struct {
	mu    sync.RWMutex
	items map[vo.ID]*warehouse.Warehouse
}

func NewWarehouseRepository() repository.WarehouseRepository {
	return &warehouseRepository{items: make(map[vo.ID]*warehouse.Warehouse)}
}

func (r *warehouseRepository) Save(ctx context.Context, w *warehouse.Warehouse) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[w.ID] = w
	return nil
}

func (r *warehouseRepository) GetByID(ctx context.Context, id vo.ID) (*warehouse.Warehouse, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	w, ok := r.items[id]
	if !ok {
		return nil, warehouse.ErrWarehouseNotFound
	}
	return w, nil
}

func (r *warehouseRepository) List(ctx context.Context, farmID vo.ID) ([]*warehouse.Warehouse, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*warehouse.Warehouse
	for _, w := range r.items {
		if w.FarmID == farmID {
			result = append(result, w)
		}
	}
	return result, nil
}

var _ repository.WarehouseRepository = (*warehouseRepository)(nil)
