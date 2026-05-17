package productionunit

import (
	"time"

	"github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	"github.com/samurenkoroma/agro-platform/internal/domain/spatial/entity/geometry"
)

type ProductionUnit struct {
	ID vo.ID

	FarmID vo.ID

	ParentID *vo.ID

	Type ProductionUnitType

	Name string

	Code *string

	Geometry *geometry.Geometry

	Capabilities []Capability

	ClimateZoneID *vo.ID

	Metadata vo.Metadata

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Aggregate struct {
	event.AggregateRoot

	Root ProductionUnit
}
