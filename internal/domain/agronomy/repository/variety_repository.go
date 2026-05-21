package repository

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/variety"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type VarietyRepository interface {
	Save(ctx context.Context, root *variety.Variety) error
	GetByID(ctx context.Context, id vo.ID) (*variety.Variety, error)
	GetByCrop(ctx context.Context, cropID vo.ID) ([]*variety.Variety, error)
	Exists(ctx context.Context, id vo.ID) (bool, error)
}
