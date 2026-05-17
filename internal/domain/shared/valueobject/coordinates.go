package valueobject

type Coordinates struct {
	X float64
	Y float64
	Z *float64
}

func NewCoordinates(
	x float64,
	y float64,
	z *float64,
) Coordinates {

	return Coordinates{
		X: x,
		Y: y,
		Z: z,
	}
}
