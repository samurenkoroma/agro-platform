package season

import (
	"context"
	"time"
)

type StatItem struct {
	Name       string  `json:"name"`
	Icon       string  `json:"icon"`
	Area       float64 `json:"area"`
	Yield      float64 `json:"yield"`
	YieldPerHa float64 `json:"yield_per_ha"`
}
type Statistics struct {
	TotalPlans     int8       `json:"totalPlans"`
	ActivePlans    int8       `json:"activePlans"`
	TotalArea      float64    `json:"totalArea"`
	CompletedPlans int8       `json:"completedPlans"`
	TotalHarvest   float64    `json:"totalHarvest"` // кг
	Crops          []StatItem `json:"crops"`
}
type Fertilizer struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}
type Resource struct {
	WaterUsed  float64      `json:"waterUsed"`
	Fertilizer []Fertilizer `json:"fertilizerUsed"`
}

type Weather struct {
	AvgTemp            float64 `json:"avgTemp"`
	TotalPrecipitation float64 `json:"totalPrecipitation"`
	SunnyDays          int     `json:"sunnyDays"`
}
type SeasonUnits struct {
	Id              string     `json:"id"`
	Name            string     `json:"name"`
	Type            string     `json:"type"`
	CropName        string     `json:"cropName"`
	VarietyName     string     `json:"varietyName"`
	Area            float64    `json:"area"`
	PlantedDate     time.Time  `json:"plantedDate"`
	HarvestDate     *time.Time `json:"harvestDate"`
	YieldEfficiency float64    `json:"yieldEfficiency"`
	Resources       Resource   `json:"resources"`
	FuelUsed        float64    `json:"fuelUsed"`
	LaborHours      float64    `json:"laborHours"`
}
type SeasonItem struct {
	Id           string        `json:"id"`
	Name         string        `json:"name"`
	StartDate    time.Time     `json:"startDate"`
	EndDate      time.Time     `json:"endDate"`
	Status       string        `json:"status"`
	Statistics   Statistics    `json:"statistics"`
	PlantingArea []SeasonUnits `json:"plantingArea"`
	Weather      Weather       `json:"weather"`
}

type SeasonFilter struct {
	OwnerId string
}

type Projection interface {
	List(context.Context, SeasonFilter) ([]SeasonItem, error)
}
