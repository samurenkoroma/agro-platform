package agronomy

import (
	"context"

	cropstage "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop_stage"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type cropStageRepository struct {
	items map[vo.ID]*cropstage.CropStage
}

func (c cropStageRepository) Save(ctx context.Context, root *cropstage.CropStage) error {
	//TODO implement me
	panic("implement me")
}

func (c cropStageRepository) GetByID(ctx context.Context, id vo.ID) (*cropstage.CropStage, error) {
	//TODO implement me
	panic("implement me")
}

func (c cropStageRepository) GetByCrop(ctx context.Context, cropID vo.ID) ([]*cropstage.CropStage, error) {
	//TODO implement me
	panic("implement me")
}

func (c cropStageRepository) Exists(ctx context.Context, id vo.ID) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewCropStageRepository() repository.CropStageRepository {
	return &cropStageRepository{items: make(map[vo.ID]*cropstage.CropStage)}
}
