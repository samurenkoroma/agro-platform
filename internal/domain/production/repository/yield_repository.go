package repository

import (
	"context"

	yieldbatch "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/yield_batch"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type YieldBatchRepository interface {
	Save(ctx context.Context, root *yieldbatch.YieldBatch) error
	GetByID(ctx context.Context, id vo.ID) (*yieldbatch.YieldBatch, error)
	GetByCycle(ctx context.Context, cycleID vo.ID) ([]*yieldbatch.YieldBatch, error)
	GetByPlant(ctx context.Context, plantID vo.ID) ([]*yieldbatch.YieldBatch, error)
	Exists(ctx context.Context, id vo.ID) (bool, error)
}
