package repository

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/disease"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type DiseaseRepository interface {
	Save(ctx context.Context, root *disease.Disease) error

	GetByID(ctx context.Context, id vo.ID) (*disease.Disease, error)

	GetByCrop(ctx context.Context, cropID vo.ID) ([]*disease.Disease, error)

	Exists(ctx context.Context, id vo.ID) (bool, error)
}
