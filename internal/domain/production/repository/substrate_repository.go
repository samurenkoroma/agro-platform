package repository

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

	substrate "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/substrate"
)

type SubstrateRepository interface {
	Save(aggregate *substrate.Aggregate) error
	GetByID(id vo.ID) (*substrate.Aggregate, error)
}
