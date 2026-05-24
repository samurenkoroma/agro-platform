package createvariety

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/commands/response"
	variety "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/variety"
	agronomy "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type CreateVarietyCommand struct {
	Name   string `json:"name" validate:"required"`
	CropID vo.ID  `json:"cropId" validate:"required"`
}

func (h *Handler) Create(ctx context.Context, payload any) (any, error) {
	cmd := payload.(*CreateVarietyCommand)

	return h.uow.Execute(ctx, providers.NewAgronomyProvider, func(provider repository.RepositoryProvider) (any, error) {
		agronomyProvider, ok := provider.(agronomy.AgronomyProvider)
		if !ok {
			return nil, repository.ErrInvalidProviderType
		}

		crop, err := agronomyProvider.Crops().GetByID(ctx, cmd.CropID)
		if err != nil {
			return nil, err
		}
		v, _ := agronomyProvider.Varieties().Exists(ctx, cmd.Name, cmd.CropID)
		if v {
			return nil, ErrVarietyAlreadyExists
		}

		root := variety.New(crop.ID, cmd.Name)

		if err := agronomyProvider.Varieties().Save(ctx, root); err != nil {
			return nil, err
		}

		h.uow.RegisterAggregate(root)

		return response.Id(root.ID), nil
	})
}
