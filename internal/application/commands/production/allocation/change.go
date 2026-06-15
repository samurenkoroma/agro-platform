package allocation

import (
	"context"
	"time"

	command "github.com/samurenkoroma/agro-platform/internal/application/commands"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/response"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	production "github.com/samurenkoroma/agro-platform/internal/domain/production/repository"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

func (h *Handler) Change(ctx context.Context, payload any) (any, error) {
	cmd, ok := payload.(*ChangeAllocationCommand)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}

	return h.uow.Execute(ctx, providers.NewProductionProvider, func(provider repository.RepositoryProvider, exec uow.Execution) (any, error) {
		productionProvider, ok := provider.(production.ProductionProvider)
		if !ok {
			return nil, repository.ErrInvalidProviderType
		}

		root, err := productionProvider.Allocation().GetByID(ctx, cmd.ID)
		if err != nil {
			return nil, err
		}

		if root == nil {
			return nil, ErrAllocationNotFound
		}

		root.ProductionUnitID = cmd.ProductionUnitID
		root.Area = cmd.Area
		root.StartedAt = cmd.StartedAt
		root.EndedAt = cmd.EndedAt
		root.UpdatedAt = time.Now()

		if err := productionProvider.Allocation().Save(ctx, root); err != nil {
			return nil, err
		}

		exec.RegisterAggregate(root)
		return response.Id(root.ID), nil
	})
}
