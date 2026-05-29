package productionunit

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Capacity struct {
	PlantCapacity     *int
	WaterVolumeLiters *float64
	AreaM2            *float64
	TrayCount         *int
	ChannelCount      *int
}

type Properties map[string]any

type ProductionUnitStatus string

const (
	Active ProductionUnitStatus = "active"

	Maintenance ProductionUnitStatus = "maintenance"

	Disabled ProductionUnitStatus = "disabled"
)

type ProductionUnit struct {
	ev.BaseAggregate
	FarmID       vo.ID
	ID           vo.ID
	ParentID     *vo.ID
	Type         ProductionUnitType
	Name         string
	Code         *string
	Description  *string
	Geometry     vo.Geometry
	Position     *vo.Position
	Capacity     *Capacity
	Capabilities []Capability
	Climate      *ClimateProfile
	Properties   Properties
	Status       ProductionUnitStatus
	Metadata     vo.Metadata
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ArchivedAt   *time.Time
}

func New(farmID vo.ID, unitType ProductionUnitType, name string) (*ProductionUnit, error) {
	if name == "" {
		return nil, ErrInvalidName
	}

	root := &ProductionUnit{
		ID:        vo.NewID(),
		FarmID:    farmID,
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
