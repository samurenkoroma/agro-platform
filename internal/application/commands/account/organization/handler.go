package organization

import (
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/jwt"
)

type OrganizationHandler struct {
	uow        uow.UnitOfWork
	jwtService *jwt.Service
}

func NewOrganizationHandler(uow uow.UnitOfWork, jwtService *jwt.Service) *OrganizationHandler {
	return &OrganizationHandler{uow: uow, jwtService: jwtService}
}
