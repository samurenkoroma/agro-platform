package repository

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

	yieldbatch "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/yield_batch"
)

type YieldRepository interface {
	Save(aggregate *yieldbatch.YieldBatch) error
	GetByID(id vo.ID) (*yieldbatch.YieldBatch, error)
	GetByHarvest(harvestID vo.ID) ([]*yieldbatch.YieldBatch, error)
}
