package productionunit

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	"github.com/samurenkoroma/agro-platform/internal/domain/spatial/entity/geometry"
)

type ProductionUnit struct {
	ev.BaseAggregate
	ID            vo.ID
	FarmID        vo.ID
	ParentID      *vo.ID
	Type          ProductionUnitType
	Name          string
	Code          *string
	Geometry      *geometry.Geometry
	Capabilities  []Capability
	ClimateZoneID *vo.ID
	Metadata      vo.Metadata
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func New(farmID vo.ID, unitType ProductionUnitType, name string) (*ProductionUnit, error) {
	if name == "" {
		return nil, ErrInvalidName
	}

	root := &ProductionUnit{
		ID:        vo.NewID(),
		FarmID:    vo.NewID(),
		Type:      unitType,
		Name:      name,
		Metadata:  vo.NewMetadata(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	root.AddEvent(NewProductionUnitCreated(root.ID))

	return root, nil
}

func (a *ProductionUnit) AttachTo(parentID vo.ID) error {
	if a.ParentID != nil {
		return ErrAlreadyHasParent
	}

	a.ParentID = &parentID

	a.AddEvent(NewProductionUnitAttached(a.ID, parentID))

	return nil
}

func (a *ProductionUnit) AddCapability(c Capability) {
	for _, v := range a.Capabilities {
		if v == c {
			return
		}
	}

	a.Capabilities = append(a.Capabilities, c)
}
