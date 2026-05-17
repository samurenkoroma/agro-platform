package valueobject

import "errors"

var ErrInvalidRange = errors.New(
	"invalid range",
)

type Range struct {
	Min float64
	Max float64
}

func NewRange(
	min float64,
	max float64,
) (Range, error) {

	if min > max {
		return Range{}, ErrInvalidRange
	}

	return Range{
		Min: min,
		Max: max,
	}, nil
}

func (r Range) Contains(
	v float64,
) bool {
	return v >= r.Min &&
		v <= r.Max
}
