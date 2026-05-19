package agronomy

import (
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type cropRepository struct {
	items map[vo.ID]*crop.Crop
}

func NewCropRepository() repository.CropRepository {
	return &cropRepository{items: make(map[vo.ID]*crop.Crop)}
}
