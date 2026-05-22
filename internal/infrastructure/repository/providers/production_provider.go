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
	//TODO implement me
	panic("implement me")
}

func (p *productionProvider) Slots() domain.SlotRepository {
	//TODO implement me
	panic("implement me")
}

func (p *productionProvider) Substrates() domain.SubstrateRepository {
	//TODO implement me
	panic("implement me")
}

func (p *productionProvider) Harvests() domain.HarvestRepository {
	//TODO implement me
	panic("implement me")
}

func (p *productionProvider) Yields() domain.YieldRepository {
	//TODO implement me
	panic("implement me")
}
