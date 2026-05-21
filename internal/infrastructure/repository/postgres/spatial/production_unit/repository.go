package productionunit

import (
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	spatial "github.com/samurenkoroma/agro-platform/internal/domain/spatial/repository"
)

type productionUnitRepository struct {
	db uow.DB
}

func New(db uow.DB) spatial.ProductionUnitRepository {
	return &productionUnitRepository{db: db}
}
