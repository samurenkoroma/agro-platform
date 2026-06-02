package planting

import (
	"context"
	"time"

	command "github.com/samurenkoroma/agro-platform/internal/application/commands"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/response"
	production "github.com/samurenkoroma/agro-platform/internal/domain/production/repository"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

func (h *Handler) Change(ctx context.Context, payload any) (any, error) {
	cmd, ok := payload.(*ChangePlantingCommand)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}

	return h.uow.Execute(ctx, providers.NewProductionProvider, func(provider repository.RepositoryProvider) (any, error) {
		productionProvider, ok := provider.(production.ProductionProvider)
		if !ok {
			return nil, repository.ErrInvalidProviderType
		}

		root, err := productionProvider.Planting().GetByID(ctx, cmd.ID)
		if err != nil {
			return nil, err
		}

		if root == nil {
			return nil, ErrPlantingNotFound
		}

		root.PlantedAt = cmd.PlantedAt
		root.Quantity = cmd.Quantity
		root.UpdatedAt = time.Now()

		if err := productionProvider.Planting().Save(ctx, root); err != nil {
			return nil, err
		}

		h.uow.RegisterAggregate(root)
		return response.Id(root.ID), nil
	})
}
