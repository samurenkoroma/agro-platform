package repository

type InventoryProvider interface {
	Items() InventoryRepository
	Movements() MovementRepository
	Warehouses() WarehouseRepository
}
