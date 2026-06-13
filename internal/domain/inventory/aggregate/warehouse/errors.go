package warehouse

import "errors"

var (
	ErrWarehouseNotFound = errors.New("warehouse not found")
	ErrWarehouseArchived = errors.New("warehouse is archived")
)
