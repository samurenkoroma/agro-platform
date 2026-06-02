package productionunit

import (
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
)

type Handler struct {
	uow uow.UnitOfWork
}

func NewProductionUnitHandler(uow uow.UnitOfWork) *Handler {
	return &Handler{uow: uow}
}
