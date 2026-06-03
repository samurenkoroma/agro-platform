package growingcycle

import (
	"context"
	"errors"

	command "github.com/samurenkoroma/agro-platform/internal/application/commands"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/response"
	"github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/allocation"
	growingcycle "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"
	production "github.com/samurenkoroma/agro-platform/internal/domain/production/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

func (h *Handler) Create(ctx context.Context, payload any) (any, error) {
	cmd, ok := payload.(*CreateCommand)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}

	orgId, ok := ctx.Value("organization_id").(string)
	if !ok {
		return nil, errors.New("organization_id is required")
	}

	return h.uow.Execute(ctx, providers.NewProductionProvider, func(provider repository.RepositoryProvider) (any, error) {
		productionProvider, ok := provider.(production.ProductionProvider)
		if !ok {
			return nil, repository.ErrInvalidProviderType
		}

		cycle := growingcycle.New(
			vo.ID(orgId), cmd.CropID,
			cmd.VarietyID, cmd.ProtocolID,
			cmd.Name, cmd.Code, cmd.Method,
			cmd.ExpectedHarvestAt)
		if err := productionProvider.GrowingCycles().Save(ctx, cycle); err != nil {
			return nil, err
		}
		h.uow.RegisterAggregate(cycle)

		alloc := allocation.New(cycle.ID, cmd.ProductionUnitID, cmd.Area, cmd.StartedAt)
		if err := productionProvider.Allocation().Save(ctx, alloc); err != nil {
			return nil, err
		}
		h.uow.RegisterAggregate(alloc)

		return response.Id(cycle.ID), nil
	})
}
