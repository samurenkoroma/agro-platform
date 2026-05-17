package valueobject

import "errors"

var ErrNegativeQuantity = errors.New(
	"negative quantity",
)

type Quantity struct {
	Value float64
	Unit  string
}

func NewQuantity(
	value float64,
	unit string,
) (Quantity, error) {

	if value < 0 {
		return Quantity{},
			ErrNegativeQuantity
	}

	return Quantity{
		Value: value,
		Unit:  unit,
	}, nil
}
