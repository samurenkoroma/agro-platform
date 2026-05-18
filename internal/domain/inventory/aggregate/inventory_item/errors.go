package inventoryitem

import "errors"

var (
	ErrNegativeAmount       = errors.New("negative amount")
	ErrInsufficientStock    = errors.New("insufficient stock")
	ErrInsufficientReserved = errors.New("insufficient reserved stock")
)
