package providers

import (
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	domain "github.com/samurenkoroma/agro-platform/internal/domain/production/repository"
	inmemory "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/production"
	postgres "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/postgres/production"
)

func (p *productionProvider) ProviderName() string {
	return ""
}

func NewProductionProvider(db uow.DB, inMemory bool) domain.ProductionProvider {
	return &productionProvider{db: db, inMemory: inMemory}
}

type productionProvider struct {
	db       uow.DB
	inMemory bool
	cycles   domain.GrowingCycleRepository
	harvests domain.HarvestBatchRepository
}

func (p *productionProvider) Harvests() domain.HarvestBatchRepository {
	//TODO implement me
	panic("implement me")
}

func (p *productionProvider) Planting() domain.PlantingRepository {
	//TODO implement me
	panic("implement me")
}

func (p *productionProvider) Allocation() domain.AllocationRepository {
	//TODO implement me
	panic("implement me")
}

func (p *productionProvider) GrowingCycles() domain.GrowingCycleRepository {
	if p.cycles != nil {
		return p.cycles
	}
	if p.inMemory {
		p.cycles = inmemory.NewGrowingCycleRepository()
	} else {
		p.cycles = postgres.NewGrowingCycleRepository(p.db)
	}
	return p.cycles
}
