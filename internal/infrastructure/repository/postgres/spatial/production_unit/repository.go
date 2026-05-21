package productionunit

import (
	spatial "github.com/samurenkoroma/agro-platform/internal/domain/spatial/repository"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type productionUnitRepository struct {
	db repository.DB
}

func New(db repository.DB) spatial.ProductionUnitRepository {
	return &productionUnitRepository{db: db}
}
