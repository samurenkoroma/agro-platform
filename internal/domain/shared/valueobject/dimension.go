package valueobject

import "errors"

var ErrInvalidDimension = errors.New(
	"invalid dimension",
)

type Dimension struct {
	Width  float64
	Length float64
	Height float64
}

func NewDimension(
	w float64,
	l float64,
	h float64,
) (Dimension, error) {

	if w < 0 ||
		l < 0 ||
		h < 0 {

		return Dimension{},
			ErrInvalidDimension
	}

	return Dimension{
		Width:  w,
		Length: l,
		Height: h,
	}, nil
}
