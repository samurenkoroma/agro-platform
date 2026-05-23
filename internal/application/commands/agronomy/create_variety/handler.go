package createvariety

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/commands/response"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	variety "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/variety"
	agronomy "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type Command struct {
	Name   string `json:"name" validate:"required"`
	CropID string `json:"cropId" validate:"required"`
}

type Handler struct {
	uow uow.UnitOfWork
}

func NewCreateVarietyHandler(uow uow.UnitOfWork) *Handler {
	return &Handler{uow: uow}
}

func (h *Handler) Handle(ctx context.Context, payload any) (any, error) {
	cmd := payload.(*Command)

	return h.uow.Execute(ctx, providers.NewAgronomyProvider, func(provider repository.RepositoryProvider) (any, error) {
		agronomyProvider, ok := provider.(agronomy.AgronomyProvider)
		if !ok {
			return nil, repository.ErrInvalidProviderType
		}

		root :=
			variety.New(vo.ID(cmd.CropID), cmd.Name)

		err := agronomyProvider.Varieties().Save(ctx, root)

		if err != nil {
			return nil, err
		}

		h.uow.RegisterAggregate(root)

		return response.Id(root.ID), nil
	},
	)
}
