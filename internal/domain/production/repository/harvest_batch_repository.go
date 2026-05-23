package repository

import (
	"context"

	harvest "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/harvest_batch"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type HarvestBatchRepository interface {
	Save(ctx context.Context, root *harvest.HarvestBatch) error
	GetByID(ctx context.Context, id vo.ID) (*harvest.HarvestBatch, error)
	GetByCycle(ctx context.Context, cycleID vo.ID) ([]*harvest.HarvestBatch, error)
	GetByProductionUnit(ctx context.Context, unitID vo.ID) ([]*harvest.HarvestBatch, error)
	Exists(ctx context.Context, id vo.ID) (bool, error)
}
