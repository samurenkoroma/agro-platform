package providers

import (
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	domain "github.com/samurenkoroma/agro-platform/internal/domain/inventory/repository"
	inmemory "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/inventory"
	postgres "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/postgres/inventory"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type inventoryProvider struct {
	db         uow.DB
	items      domain.InventoryRepository
	movements  domain.MovementRepository
	warehouses domain.WarehouseRepository
}

func NewInventoryProvider(db uow.DB) repository.RepositoryProvider {
	return &inventoryProvider{db: db}
}

func (p *inventoryProvider) ProviderName() string { return "inventory" }

func (p *inventoryProvider) Items() domain.InventoryRepository {
	if p.items == nil {
		p.items = postgres.NewInventoryRepository(p.db)
	}
	return p.items
}

func (p *inventoryProvider) Movements() domain.MovementRepository {
	if p.movements == nil {
		p.movements = postgres.NewMovementRepository(p.db)
	}
	return p.movements
}

func (p *inventoryProvider) Warehouses() domain.WarehouseRepository {
	if p.warehouses == nil {
		p.warehouses = postgres.NewWarehouseRepository(p.db)
	}
	return p.warehouses
}

func NewInventoryProviderInMemory(_ uow.DB) repository.RepositoryProvider {
	return &inventoryProvider{
		items:      inmemory.NewInventoryRepository(),
		movements:  inmemory.NewMovementRepository(),
		warehouses: inmemory.NewWarehouseRepository(),
	}
}

var _ domain.InventoryProvider = (*inventoryProvider)(nil)
