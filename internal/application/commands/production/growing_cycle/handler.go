package growingcycle

import (
	unitOfWork "github.com/samurenkoroma/agro-platform/internal/application/uow"
)

type Handler struct {
	uow unitOfWork.UnitOfWork
}

func NewGrowingCycleHandler(uow unitOfWork.UnitOfWork) *Handler {
	return &Handler{uow: uow}
}
