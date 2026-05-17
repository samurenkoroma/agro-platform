package valueobject

import (
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidID = errors.New("invalid id")

type ID string

func NewID() ID {
	return ID(uuid.NewString())
}

func ParseID(v string) (ID, error) {
	if v == "" {
		return "", ErrInvalidID
	}

	_, err := uuid.Parse(v)
	if err != nil {
		return "", ErrInvalidID
	}

	return ID(v), nil
}

func (i ID) String() string {
	return string(i)
}

func (i ID) IsZero() bool {
	return i == ""
}
