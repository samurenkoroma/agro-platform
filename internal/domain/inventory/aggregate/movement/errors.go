package movement

import "errors"

var (
	ErrMovementNotFound = errors.New("movement not found")
	ErrInvalidQuantity  = errors.New("quantity must be positive")
)
