package variety

import (
	"context"
)

type ListFilter struct {
	CropKey string
}

type Detail struct {
	ID               string `json:"id"`
	Name             string
	SpeciesKey       string
	SpeciesName      string
	BaseTemperature  float64
	MaxTemperature   float64
	DaysToMaturity   int
	PhenophaseGDD    []Phenophase
	WaterRequirement WaterRequirement
	LightRequirement LightRequirement
	YieldPotential   string
	GrowingTypes     []string
	Description      string
	Image            string
}
type Phenophase struct {
	Code        string
	Name        string
	GddRequired int
	Description string
	IsCritical  bool
}
type WaterRequirement struct {
	DailyNeedMin   int
	DailyNeedOpt   int
	CriticalPhases []string
}
type LightRequirement struct {
	PpfdMin         int
	PpfdOpt         int
	DayLengthMin    int
	DayLengthOpt    int
	PhotoperiodType string
	CriticalPhases  []string
}

type ListItem struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	CropId   string `json:"cropId"`
	CropName string `json:"cropName"`
}

type Projection interface {
	Get(context.Context, string) (*Detail, error)
	List(context.Context, ListFilter) ([]ListItem, error)
}
