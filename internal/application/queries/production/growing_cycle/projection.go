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

type Projection interface {
	Get(ctx context.Context, id vo.ID) (*DTO, error)
	List(ctx context.Context, ownerId vo.ID) ([]*DTO, error)
}
