package repository

import "github.com/samurenkoroma/agro-platform/internal/domain/inventory/repository"

type inventoryRepository struct{}

func NewInventoryRepository() repository.InventoryRepository {
	return &inventoryRepository{}
}
