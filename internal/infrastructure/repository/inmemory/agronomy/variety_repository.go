package agronomy

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/variety"
	agronomyrepo "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type varietyRepository struct {
	items map[vo.ID]*variety.Variety
}

func (v varietyRepository) Save(ctx context.Context, root *variety.Variety) error {
	//TODO implement me
	panic("implement me")
}

func (v varietyRepository) GetByID(ctx context.Context, id vo.ID) (*variety.Variety, error) {
	//TODO implement me
	panic("implement me")
}

func (v varietyRepository) GetByCrop(ctx context.Context, cropID vo.ID) ([]*variety.Variety, error) {
	//TODO implement me
	panic("implement me")
}

func (v varietyRepository) Exists(ctx context.Context, name string, cropId vo.ID) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewVarietyRepository() agronomyrepo.VarietyRepository {
	return &varietyRepository{items: make(map[vo.ID]*variety.Variety)}
}
