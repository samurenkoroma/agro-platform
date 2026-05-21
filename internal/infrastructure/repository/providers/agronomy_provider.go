package providers

import (
	domain "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	inmemory "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/agronomy"
	postgres "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/postgres/agronomy"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

func (p *agronomyProvider) ProviderName() string {
	return "agronomy"
}

func NewAgronomyProvider(db repository.DB, inMemory bool) domain.AgronomyProvider {
	return &agronomyProvider{db: db, inMemory: inMemory}
}

type agronomyProvider struct {
	db             repository.DB
	inMemory       bool
	crops          domain.CropRepository
	varieties      domain.VarietyRepository
	protocols      domain.CropProtocolRepository
	diseases       domain.DiseaseRepository
	stressProfiles domain.StressRepository
}

func (p *agronomyProvider) Crops() domain.CropRepository {
	if p.crops != nil {
		return p.crops
	}
	if p.inMemory {
		p.crops = inmemory.NewCropRepository()
	} else {
		p.crops = postgres.NewCropRepository(p.db)
	}
	return p.crops
}

func (p *agronomyProvider) Varieties() domain.VarietyRepository {
	if p.varieties != nil {
		return p.varieties
	}
	if p.inMemory {
		p.varieties = inmemory.NewVarietyRepository()
	} else {
		p.varieties = postgres.NewVarietyRepository(p.db)
	}

	return p.varieties
}

func (p *agronomyProvider) Protocols() domain.CropProtocolRepository {
	if p.protocols != nil {
		return p.protocols
	}
	if p.inMemory {
		p.protocols = inmemory.NewProtocolRepository()
	} else {
		p.protocols = postgres.NewProtocolRepository(p.db)
	}
	return p.protocols
}

func (p *agronomyProvider) Diseases() domain.DiseaseRepository {
	//TODO implement me
	panic("implement me")
}

func (p *agronomyProvider) StressProfiles() domain.StressRepository {
	//TODO implement me
	panic("implement me")
}
