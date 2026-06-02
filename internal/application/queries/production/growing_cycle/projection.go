package growingcycle

import (
	"context"
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type DTO struct {
	ID          vo.ID  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	CropName    string `json:"cropName"`
	VarietyName string `json:"varietyName"`
	Status      string `json:"status"`
	Stage       string `json:"stage"`

	ExpectedHarvestAt *time.Time `json:"expectedHarvestAt"`
	CreatedAt         time.Time  `json:"createdAt"`
}

type SummaryDTO struct {
	ID                vo.ID
	Name              string
	Status            string
	AllocatedArea     float64
	PlantedQuantity   float64
	HarvestedQuantity float64
	ExpectedHarvestAt *time.Time
}

type Projection interface {
	Get(ctx context.Context, id vo.ID) (*DTO, error)
	List(ctx context.Context, ownerId vo.ID) ([]*DTO, error)
	Summary(ctx context.Context, ownerId vo.ID, cycleId vo.ID) (*SummaryDTO, error)
}
