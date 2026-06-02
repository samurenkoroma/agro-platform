package growingcycle

import (
	"context"
	"errors"
	"time"

	command "github.com/samurenkoroma/agro-platform/internal/application/commands"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/response"
	production "github.com/samurenkoroma/agro-platform/internal/domain/production/repository"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

func (h *Handler) Update(ctx context.Context, payload any) (any, error) {
	cmd, ok := payload.(*UpdateCommand)
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

		cycle, err := productionProvider.GrowingCycles().GetByID(ctx, cmd.ID)
		if err != nil || cycle.FarmID.String() != orgId {
			return nil, ErrGrowingCycleNotFound
		}

		cycle.CropID = cmd.CropID
		if cmd.VarietyID != nil {
			cycle.VarietyID = cmd.VarietyID
		}
		if cmd.ProtocolID != nil {
			cycle.ProtocolID = cmd.ProtocolID
		}

		cycle.Name = cmd.Name
		cycle.Code = cmd.Code

		cycle.Method = cmd.Method

		cycle.Status = cmd.Status
		cycle.Stage = cmd.Stage

		cycle.ExpectedHarvestAt = cmd.ExpectedHarvestAt
		cycle.UpdatedAt = time.Now()

		if err := productionProvider.GrowingCycles().Save(ctx, cycle); err != nil {
			return nil, err
		}

		h.uow.RegisterAggregate(cycle)
		return response.Id(cycle.ID), nil
	})
}
