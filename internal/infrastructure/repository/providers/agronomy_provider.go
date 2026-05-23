package providers

import (
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	domain "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	inmemory "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/agronomy"
	postgres "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/postgres/agronomy"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

func (p *agronomyProvider) ProviderName() string {
	return "agronomy"
}

func NewAgronomyProvider(db uow.DB) repository.RepositoryProvider {
	return &agronomyProvider{db: db, inMemory: false}
}

type agronomyProvider struct {
	db             uow.DB
	inMemory       bool
	crops          domain.CropRepository
	cropsStages    domain.CropStageRepository
	varieties      domain.VarietyRepository
	protocols      domain.CropProtocolRepository
	diseases       domain.DiseaseRepository
	stressProfiles domain.StressRepository
}

func (p *agronomyProvider) CropsStages() domain.CropStageRepository {
	if p.cropsStages != nil {
		return p.cropsStages
	}
	if p.inMemory {
		p.cropsStages = inmemory.NewCropStageRepository()
	} else {
		p.cropsStages = postgres.NewCropStageRepository(p.db)
	}
	return p.cropsStages
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
	if p.diseases != nil {
		return p.diseases
	}
	if p.inMemory {
		p.diseases = inmemory.NewDiseaseRepository()
	} else {
		p.diseases = postgres.NewDiseaseRepository(p.db)
	}
	return p.diseases
}

func (p *agronomyProvider) StressProfiles() domain.StressRepository {
	if p.stressProfiles != nil {
		return p.stressProfiles
	}
	if p.inMemory {
		p.stressProfiles = inmemory.NewStressRepository()
	} else {
		p.stressProfiles = postgres.NewStressRepository(p.db)
	}
	return p.stressProfiles
}
