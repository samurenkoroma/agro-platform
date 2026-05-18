package substrate

import "errors"

var (
	ErrAlreadyDisposed = errors.New("substrate disposed")
	ErrInvalidVolume   = errors.New("invalid volume")
	ErrNotReusable     = errors.New("substrate not reusable")
)
