package productionunitsnapshot

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
)

type ProductionUnitSnapshot struct {
	ID vo.ID

	SnapshotID vo.ID

	OriginalUnitID vo.ID

	Type pu.ProductionUnitType

	Name string

	ParentID *vo.ID

	Capabilities []pu.Capability

	Metadata vo.Metadata

	CreatedAt time.Time
}
