package agronomy

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type cropRepository struct {
	items map[vo.ID]*crop.Crop
}

func (c cropRepository) Save(ctx context.Context, entity *crop.Crop) error {
	//TODO implement me
	panic("implement me")
}

func (c cropRepository) GetByID(ctx context.Context, id vo.ID) (*crop.Crop, error) {
	//TODO implement me
	panic("implement me")
}

func (c cropRepository) Exists(ctx context.Context, id vo.ID) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c cropRepository) GetAll(ctx context.Context) ([]*crop.Crop, error) {
	//TODO implement me
	panic("implement me")
}

func NewCropRepository() repository.CropRepository {
	return &cropRepository{items: make(map[vo.ID]*crop.Crop)}
}
