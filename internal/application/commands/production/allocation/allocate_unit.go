package allocation

import (
	"context"

	command "github.com/samurenkoroma/agro-platform/internal/application/commands"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/response"
	"github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/allocation"
	production "github.com/samurenkoroma/agro-platform/internal/domain/production/repository"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

func (h *Handler) AllocateProductionUnit(ctx context.Context, payload any) (any, error) {
	cmd, ok := payload.(*AllocateProductionUnitCommand)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}

	return h.uow.Execute(ctx, providers.NewProductionProvider, func(provider repository.RepositoryProvider) (any, error) {
		productionProvider, ok := provider.(production.ProductionProvider)
		if !ok {
			return nil, repository.ErrInvalidProviderType
		}

		item := allocation.New(
			cmd.CycleID,
			cmd.ProductionUnitID,
			cmd.Area,
			cmd.StartedAt,
		)

		if err := productionProvider.Allocation().Save(ctx, item); err != nil {
			return nil, err
		}

		h.uow.RegisterAggregate(item)
		return response.Id(item.ID), nil
	})
}
