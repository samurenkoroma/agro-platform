package valueobject

import "errors"

var ErrInvalidArea = errors.New("invalid area")

type Area struct {
	SquareMeters float64
}

func NewArea(v float64) (Area, error) {

	if v < 0 {
		return Area{},
			ErrInvalidArea
	}

	return Area{
		SquareMeters: v,
	}, nil
}
