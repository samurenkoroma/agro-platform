package valueobject

type GeometryType string

const (
	Rect    GeometryType = "RECT"
	Polygon GeometryType = "POLYGON"
	Point   GeometryType = "POINT"
)

type Geometry struct {
	Type GeometryType

	Dimension *Dimension

	Position *Coordinates

	Polygon [][]float64
}
