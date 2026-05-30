package productionunit

import "errors"

var (
	ErrInvalidHierarchy = errors.New("invalid hierarchy")
	ErrAlreadyHasParent = errors.New("unit already has parent")
	ErrInvalidCode      = errors.New("invalid code")
)
