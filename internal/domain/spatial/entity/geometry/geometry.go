package geometry

import vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

type GeometryType string

const (
	Rect    GeometryType = "RECT"
	Polygon GeometryType = "POLYGON"
	Point   GeometryType = "POINT"
)

type Geometry struct {
	Type GeometryType

	Dimension *vo.Dimension

	Position *vo.Coordinates

	Polygon [][]float64
}
