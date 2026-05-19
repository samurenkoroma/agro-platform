package repository

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

	harvest "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/harvest_batch"
)

type HarvestRepository interface {
	Save(aggregate *harvest.HarvestBatch) error
	GetByID(id vo.ID) (*harvest.HarvestBatch, error)
	GetByCycle(cycleID vo.ID) ([]*harvest.HarvestBatch, error)
}
