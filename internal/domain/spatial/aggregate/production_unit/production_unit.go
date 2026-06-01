package productionunit

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type ProductionUnitStatus string

const (
	Active      ProductionUnitStatus = "active"
	Maintenance ProductionUnitStatus = "maintenance"
	Disabled    ProductionUnitStatus = "disabled"
)

type ProductionUnit struct {
	ev.BaseAggregate
	OwnerID    vo.ID
	ID         vo.ID
	ParentID   *vo.ID
	Type       ProductionUnitType
	Code       string
	Geometry   vo.Geometry
	Properties *Properties
	Status     ProductionUnitStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ArchivedAt *time.Time
}

func New(
	ownerID vo.ID,
	ParentId *vo.ID,
	unitType ProductionUnitType,
	status ProductionUnitStatus,
	code string,
) *ProductionUnit {
	now := time.Now()
	root := &ProductionUnit{
		ID:         vo.NewID(),
		ParentID:   ParentId,
		OwnerID:    ownerID,
		Code:       code,
		Type:       unitType,
		Status:     status,
		Properties: NewProps(code, ""),
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	root.AddEvent(NewProductionUnitCreated(root.ID))

	return root
}

func (obj *ProductionUnit) AddDimensions(dim *Dimensions) {
	obj.Properties.Dimensions = dim
}
