package agronomy

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/disease"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type diseaseRepository struct{}

func (d diseaseRepository) Save(ctx context.Context, root *disease.Disease) error {
	//TODO implement me
	panic("implement me")
}

func (d diseaseRepository) GetByID(ctx context.Context, id vo.ID) (*disease.Disease, error) {
	//TODO implement me
	panic("implement me")
}

func (d diseaseRepository) GetAll(ctx context.Context) ([]*disease.Disease, error) {
	//TODO implement me
	panic("implement me")
}

func (d diseaseRepository) Exists(ctx context.Context, id vo.ID) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewDiseaseRepository() repository.DiseaseRepository {
	return &diseaseRepository{}
}
