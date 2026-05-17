package spatial

import "github.com/samurenkoroma/agro-platform/internal/domain/shared"

// ProductionUnit — главный Aggregate Root в Spatial Context
type ProductionUnit struct {
	ID       shared.ID
	Name     string
	Type     string // Field, Greenhouse, Orchard и т.д.
	Geometry string // GeoJSON или PostGIS
	Area     float64
	// ... другие атрибуты
}

func NewProductionUnit(name string, unitType string) *ProductionUnit {
	return &ProductionUnit{
		ID:   shared.ID("pu_" + name), // временно
		Name: name,
		Type: unitType,
	}
}
