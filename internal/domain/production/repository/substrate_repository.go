package repository

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

	substrate "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/substrate"
)

type SubstrateRepository interface {
	Save(aggregate *substrate.Substrate) error
	GetByID(id vo.ID) (*substrate.Substrate, error)
}
