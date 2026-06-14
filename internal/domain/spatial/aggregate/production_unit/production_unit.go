package productionunit

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type ProductionUnitStatus string

const (
	Growing     ProductionUnitStatus = "growing"
	Preparation ProductionUnitStatus = "preparation"
	Empty       ProductionUnitStatus = "empty"
)

type ProductionUnit struct {
	ev.BaseAggregate
	OwnerID    vo.ID
	ID         vo.ID
	ParentID   *vo.ID
	Type       ProductionUnitType
	Code       string
	Area       float64
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
	code string,
	name *string,
) *ProductionUnit {
	now := time.Now()
	root := &ProductionUnit{
		ID:         vo.NewID(),
		ParentID:   ParentId,
		OwnerID:    ownerID,
		Code:       code,
		Type:       unitType,
		Status:     Empty,
		Properties: NewProps(*name, ""),
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	root.AddEvent(NewProductionUnitCreated(root.ID))

	return root
}

func (obj *ProductionUnit) AddDimensions(dim *Dimensions) {
	obj.Properties.Dimensions = dim
	if dim.Width != nil && dim.Length != nil {
		w := *dim.Width
		l := *dim.Length
		obj.Area = (w * l) / 10000
	}
}

func (obj *ProductionUnit) Occupy() {
	obj.Status = Growing
	obj.UpdatedAt = time.Now()
	//obj.AddEvent(NewProductionUnitOccupied(obj.ID))
}

func (obj *ProductionUnit) Release() {
	obj.Status = Empty
	obj.UpdatedAt = time.Now()
	obj.AddEvent(NewProductionUnitReleased(obj.ID))
}

func (obj *ProductionUnit) SetPreparation() {
	obj.Status = Preparation
	obj.UpdatedAt = time.Now()
	obj.AddEvent(NewProductionUnitInPreparation(obj.ID))
}
