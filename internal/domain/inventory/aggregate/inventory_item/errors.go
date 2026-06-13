package inventoryitem

import "errors"

var (
	ErrItemNotFound         = errors.New("inventory item not found")
	ErrNegativeAmount       = errors.New("amount must be positive")
	ErrInsufficientStock    = errors.New("insufficient available stock")
	ErrInsufficientReserved = errors.New("insufficient reserved stock")
)
