package item

import unitOfWork "github.com/samurenkoroma/agro-platform/internal/application/uow"

type Handler struct{ uow unitOfWork.UnitOfWork }

func NewItemHandler(uow unitOfWork.UnitOfWork) *Handler { return &Handler{uow: uow} }
