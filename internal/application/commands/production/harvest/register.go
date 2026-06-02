package harvest

import (
	"context"

	command "github.com/samurenkoroma/agro-platform/internal/application/commands"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/response"
	"github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/harvest_batch"
	production "github.com/samurenkoroma/agro-platform/internal/domain/production/repository"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

func (h *Handler) Register(ctx context.Context, payload any) (any, error) {
	cmd, ok := payload.(*RegisterHarvestCommand)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}

	return h.uow.Execute(ctx, providers.NewProductionProvider, func(provider repository.RepositoryProvider) (any, error) {
		productionProvider, ok := provider.(production.ProductionProvider)
		if !ok {
			return nil, repository.ErrInvalidProviderType
		}

		root := harvestbatch.New(
			cmd.CycleID,
			cmd.HarvestAt,
			cmd.Quantity,
		)

		if err := productionProvider.Harvests().Save(ctx, root); err != nil {
			return nil, err
		}

		h.uow.RegisterAggregate(root)
		return response.Id(root.ID), nil
	})
}
