package disease

import "errors"

var (
	ErrInvalidSeverity = errors.New("invalid severity")
	ErrAlreadyResolved = errors.New("disease already resolved")
	ErrSymptomEmpty    = errors.New("symptom description empty")
)
