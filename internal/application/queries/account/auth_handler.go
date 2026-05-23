package account

import (
	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/jwt"
)

type UserHandler struct {
	uow        uow.UnitOfWork
	jwtService *jwt.Service
}

func NewUserHandler(uow uow.UnitOfWork, jwtService *jwt.Service) queries.Handler {
	return &UserHandler{
		uow:        uow,
		jwtService: jwtService,
	}
}
