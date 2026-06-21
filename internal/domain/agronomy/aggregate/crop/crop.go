package crop

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Crop struct {
	ev.BaseAggregate
	ID                vo.ID
	Name              string
	Family            string
	Category          CropCategory
	DefaultProtocolID *vo.ID

	Agronomy AgronomyProfile
	Metadata vo.Metadata

	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	StageEmergence     = "EMERGENCE"
	StageVegetative    = "VEGETATIVE"
	StageFlowering     = "FLOWERING"
	StageFruitSet      = "FRUIT_SET"
	StageFruitFill     = "FRUIT_FILL"
	StageHeadFormation = "HEAD_FORMATION"
	StageRootFill      = "ROOT_FILL"
	StageHarvest       = "HARVEST"
)

type AgronomyProfile struct {
	ThermalRequirements ThermalRequirements
	Tolerance           Tolerance
	WaterRequirement    WaterRequirement
	LightRequirement    LightRequirement
	PhenophaseGDD       []PhenophaseGDD
	Maturity            vo.Maturity
}

type ThermalRequirements struct {
	BaseTemperature  float64
	UpperTemperature float64
}

type Tolerance struct {
	TemperatureMin int8
	TemperatureMax int8
	HumidityMin    int8
	HumidityMax    int8
}

type WaterRequirement struct {
	DailyNeedMin float64
	DailyNeedOpt float64
}

type LightRequirement struct {
	PPFDMin         int
	PPFDOpt         int
	DayLengthMin    int8
	DayLengthOpt    int8
	PhotoperiodType string
}

type PhenophaseGDD struct {
	Code        string  // "BBCH-10"
	Name        string  // "Первый настоящий лист"
	GDDRequired float64 // накопленное GDD для достижения
	Description string  // описание фазы
	IsCritical  bool    // критическая фаза?
}

func New(name string, category CropCategory, family string, agronomy AgronomyProfile) *Crop {
	now := time.Now()

	root := &Crop{
		ID:        vo.NewID(),
		Family:    family,
		Category:  category,
		Name:      name,
		Agronomy:  agronomy,
		Metadata:  vo.NewMetadata(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	root.AddEvent(NewCropCreated(root.ID))

	return root
}

func (a *Crop) Rename(name string) {
	a.Name = name
	a.UpdatedAt = time.Now()

	a.AddEvent(NewCropRenamed(a.ID))
}

func (a *Crop) AssignProtocol(id vo.ID) {
	a.DefaultProtocolID = &id
	a.UpdatedAt = time.Now()

	a.AddEvent(NewProtocolAssigned(a.ID, id))
}
