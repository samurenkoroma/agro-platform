package production

import (
	"context"

	harvest "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/harvest_batch"
	"github.com/samurenkoroma/agro-platform/internal/domain/production/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type harvestBatchRepository struct {
}

func (h harvestBatchRepository) Save(ctx context.Context, root *harvest.HarvestBatch) error {
	//TODO implement me
	panic("implement me")
}

func (h harvestBatchRepository) GetByID(ctx context.Context, id vo.ID) (*harvest.HarvestBatch, error) {
	//TODO implement me
	panic("implement me")
}

func (h harvestBatchRepository) GetByCycle(ctx context.Context, cycleID vo.ID) ([]*harvest.HarvestBatch, error) {
	//TODO implement me
	panic("implement me")
}

func (h harvestBatchRepository) GetByProductionUnit(ctx context.Context, unitID vo.ID) ([]*harvest.HarvestBatch, error) {
	//TODO implement me
	panic("implement me")
}

func (h harvestBatchRepository) Exists(ctx context.Context, id vo.ID) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewHarvestBatchRepository() repository.HarvestBatchRepository {
	return &harvestBatchRepository{}
}
