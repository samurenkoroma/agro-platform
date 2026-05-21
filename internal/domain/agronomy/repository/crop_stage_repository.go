package repository

import (
	"context"

	stage "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop_stage"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type CropStageRepository interface {
	Save(ctx context.Context, root *stage.CropStage) error
	GetByID(ctx context.Context, id vo.ID) (*stage.CropStage, error)
	GetByCrop(ctx context.Context, cropID vo.ID) ([]*stage.CropStage, error)
	Exists(ctx context.Context, id vo.ID) (bool, error)
}
