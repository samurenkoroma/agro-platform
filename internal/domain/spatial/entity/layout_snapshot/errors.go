package layoutsnapshot

import "errors"

var (
	ErrInvalidVersion = errors.New("invalid version")
	ErrDuplicateUnit  = errors.New("duplicate unit")
)
