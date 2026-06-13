package inventory

import (
	"context"
	"sync"

	inventoryitem "github.com/samurenkoroma/agro-platform/internal/domain/inventory/aggregate/inventory_item"
	"github.com/samurenkoroma/agro-platform/internal/domain/inventory/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type inventoryRepository struct {
	mu    sync.RWMutex
	items map[vo.ID]*inventoryitem.Item
}

func NewInventoryRepository() repository.InventoryRepository {
	return &inventoryRepository{items: make(map[vo.ID]*inventoryitem.Item)}
}

func (r *inventoryRepository) Save(ctx context.Context, item *inventoryitem.Item) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[item.ID] = item
	return nil
}

func (r *inventoryRepository) GetByID(ctx context.Context, id vo.ID) (*inventoryitem.Item, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	item, ok := r.items[id]
	if !ok {
		return nil, inventoryitem.ErrItemNotFound
	}
	return item, nil
}

func (r *inventoryRepository) List(ctx context.Context, filter repository.ItemFilter) ([]*inventoryitem.Item, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*inventoryitem.Item
	for _, item := range r.items {
		if item.FarmID != filter.FarmID {
			continue
		}
		if filter.WarehouseID != nil && (item.WarehouseID == nil || *item.WarehouseID != *filter.WarehouseID) {
			continue
		}
		if filter.Type != nil && item.Type != *filter.Type {
			continue
		}
		if !filter.Archived && item.ArchivedAt != nil {
			continue
		}
		result = append(result, item)
	}
	return result, nil
}

var _ repository.InventoryRepository = (*inventoryRepository)(nil)
