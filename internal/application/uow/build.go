package uow

import (
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type ProviderName string

const (
	Spatial ProviderName = "spatial"
)

type ProviderDeps struct {
	Name     ProviderName
	InMemory bool
}

func Deps(name ProviderName, inMemory bool) ProviderDeps {
	return ProviderDeps{
		Name:     name,
		InMemory: inMemory,
	}
}

func InMemoryDeps(name ProviderName) ProviderDeps {
	return ProviderDeps{
		Name:     name,
		InMemory: true,
	}
}

func BuildProvider(db repository.DB, deps ProviderDeps) repository.RepositoryProvider {
	switch deps.Name {
	case Spatial:
		return providers.NewSpatialProvider(db, deps.InMemory)

	default:
		return nil
	}
}
