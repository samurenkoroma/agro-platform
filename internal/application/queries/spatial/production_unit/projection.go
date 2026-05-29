package productionunit

import (
	"context"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type CapacityDTO struct {
	PlantCapacity     *int     `json:"plantCapacity"`
	WaterVolumeLiters *float64 `json:"waterVolumeLiters"`
	AreaM2            *float64 `json:"areaM2"`
	TrayCount         *int     `json:"trayCount"`
	ChannelCount      *int     `json:"channelCount"`
}

type ClimateDTO struct {
	TemperatureMin *float64 `json:"temperatureMin"`
	TemperatureMax *float64 `json:"temperatureMax"`
	HumidityMin    *float64 `json:"humidityMin"`
	HumidityMax    *float64 `json:"humidityMax"`
	CO2Min         *float64 `json:"co2Min"`
	CO2Max         *float64 `json:"co2Max"`
	LightPPFDMin   *float64 `json:"lightPpfdMin"`
	LightPPFDMax   *float64 `json:"lightPpfdMax"`
}

type TreeNode struct {
	ID       vo.ID      `json:"id"`
	ParentID *vo.ID     `json:"parentId"`
	Type     string     `json:"type"`
	Name     string     `json:"name"`
	Status   string     `json:"status"`
	Children []TreeNode `json:"children"`
}

type DTO struct {
	ID          vo.ID          `json:"id"`
	ParentID    *vo.ID         `json:"parentId"`
	Type        string         `json:"type"`
	Status      string         `json:"status"`
	Name        string         `json:"name"`
	Code        *string        `json:"code"`
	Description *string        `json:"description"`
	Geometry    *vo.Geometry   `json:"geometry"`
	Position    *vo.Position   `json:"position"`
	Capacity    *CapacityDTO   `json:"capacity"`
	Climate     *ClimateDTO    `json:"climate"`
	Properties  map[string]any `json:"properties"`
	Metadata    vo.Metadata    `json:"metadata"`
}

type Projection interface {
	Get(ctx context.Context, id vo.ID) (*DTO, error)
	ListRoots(ctx context.Context) ([]DTO, error)
	ListChildren(ctx context.Context, parentID vo.ID) ([]DTO, error)
	Tree(ctx context.Context, rootID *vo.ID) ([]TreeNode, error)
}
