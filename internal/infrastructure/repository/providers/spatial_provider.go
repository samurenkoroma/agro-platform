package providers

import (
	domain "github.com/samurenkoroma/agro-platform/internal/domain/spatial/repository"
	inmemory "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/spatial"
	postgres "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/postgres/spatial/production_unit"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type spatialProvider struct {
	db       repository.DB
	inMemory bool
	units    domain.ProductionUnitRepository
}

func (p *spatialProvider) ProviderName() string {
	return "spatial"
}

func NewSpatialProvider(db repository.DB, inMemory bool) domain.SpatialProvider {
	return &spatialProvider{db: db, inMemory: inMemory}
}

func (p *spatialProvider) Units() domain.ProductionUnitRepository {
	if p.units != nil {
		return p.units
	}

	if p.inMemory {
		p.units = inmemory.NewProductionUnitRepository()
	} else {
		p.units = postgres.New(p.db)
	}
	return p.units
}
