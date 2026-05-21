package providers

import (
	"github.com/samurenkoroma/agro-platform/internal/application/provider"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

const (
	Spatial  provider.ProviderName = "spatial"
	Agronomy provider.ProviderName = "agronomy"
)

func BuildProvider(db uow.DB, deps provider.ProviderDeps) repository.RepositoryProvider {
	switch deps.Name {
	case Spatial:
		return NewSpatialProvider(db, deps.InMemory)
	case Agronomy:
		return NewAgronomyProvider(db, deps.InMemory)

	default:
		return nil
	}
}
