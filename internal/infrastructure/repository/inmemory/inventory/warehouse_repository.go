package repository

import "github.com/samurenkoroma/agro-platform/internal/domain/inventory/repository"

type warehouseRepository struct{}

func NewWarehouseRepository() repository.WarehouseRepository {
	return &warehouseRepository{}
}
