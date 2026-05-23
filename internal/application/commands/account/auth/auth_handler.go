package auth

import (
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/jwt"
)

// AuthHandler обрабатывает auth запросы (без CQRS)
type AuthHandler struct {
	uow        uow.UnitOfWork
	jwtService *jwt.Service
}

// NewAuthHandler создает новый AuthHandler
func NewAuthHandler(uow uow.UnitOfWork, jwtService *jwt.Service) *AuthHandler {
	return &AuthHandler{
		uow:        uow,
		jwtService: jwtService,
	}
}
