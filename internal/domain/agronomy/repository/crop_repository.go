package repository

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type CropRepository interface {
	Save(ctx context.Context, entity *crop.Crop) error
	GetByID(ctx context.Context, id vo.ID) (*crop.Crop, error)
	Exists(ctx context.Context, key string) (bool, error)
	GetAll(ctx context.Context) ([]*crop.Crop, error)
}
