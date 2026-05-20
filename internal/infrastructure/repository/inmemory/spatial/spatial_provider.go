package spatial

import (
	"database/sql"

	spatial "github.com/samurenkoroma/agro-platform/internal/domain/spatial/repository"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type InMemorySpatialProvider struct {
	tx *sql.Tx

	// Кеш репозиториев
	units spatial.ProductionUnitRepository
}

func (p *InMemorySpatialProvider) ProviderName() string {
	return "spatial"
}

// Проверяем, что FarmProvider реализует интерфейс RepositoryProvider
var _ repository.RepositoryProvider = (*InMemorySpatialProvider)(nil)

func NewInMemorySpatialProvider(tx *sql.Tx) repository.RepositoryProvider {
	return &InMemorySpatialProvider{}
}

func (p *InMemorySpatialProvider) Units() spatial.ProductionUnitRepository {
	if p.units == nil {
		p.units = NewProductionUnitRepository()
	}
	return p.units
}
