package production

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	yieldbatch "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/yield_batch"
	"github.com/samurenkoroma/agro-platform/internal/domain/production/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type yieldRepository struct {
	db uow.DB
}

func (y yieldRepository) Save(ctx context.Context, root *yieldbatch.YieldBatch) error {
	//TODO implement me
	panic("implement me")
}

func (y yieldRepository) GetByID(ctx context.Context, id vo.ID) (*yieldbatch.YieldBatch, error) {
	//TODO implement me
	panic("implement me")
}

func (y yieldRepository) GetByCycle(ctx context.Context, cycleID vo.ID) ([]*yieldbatch.YieldBatch, error) {
	//TODO implement me
	panic("implement me")
}

func (y yieldRepository) GetByPlant(ctx context.Context, plantID vo.ID) ([]*yieldbatch.YieldBatch, error) {
	//TODO implement me
	panic("implement me")
}

func (y yieldRepository) Exists(ctx context.Context, id vo.ID) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewYieldBatchRepository() repository.YieldBatchRepository {
	return &yieldRepository{}
}
