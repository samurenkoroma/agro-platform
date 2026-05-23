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
	plants   domain.PlantRepository
	harvests domain.HarvestBatchRepository
	yields   domain.YieldBatchRepository
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

func (p *productionProvider) Plants() domain.PlantRepository {
	if p.plants != nil {
		return p.plants
	}
	if p.inMemory {
		p.plants = inmemory.NewPlantRepository()
	} else {
		p.plants = postgres.NewPlantRepository(p.db)
	}
	return p.plants
}

func (p *productionProvider) Slots() domain.SlotRepository {
	//TODO implement me
	panic("implement me")
}

func (p *productionProvider) Substrates() domain.SubstrateRepository {
	//TODO implement me
	panic("implement me")
}

func (p *productionProvider) Harvests() domain.HarvestBatchRepository {
	if p.harvests != nil {
		return p.harvests
	}
	if p.inMemory {
		p.harvests = inmemory.NewHarvestBatchRepository()
	} else {
		p.harvests = postgres.NewHarvestBatchRepository(p.db)
	}
	return p.harvests
}

func (p *productionProvider) Yields() domain.YieldBatchRepository {
	if p.yields != nil {
		return p.yields
	}
	if p.inMemory {
		p.yields = inmemory.NewYieldBatchRepository()
	} else {
		p.yields = postgres.NewYieldBatchRepository(p.db)
	}
	return p.yields
}
