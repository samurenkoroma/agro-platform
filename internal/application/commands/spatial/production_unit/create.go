package productionunit

import (
	"context"
	"errors"
	"fmt"

	command "github.com/samurenkoroma/agro-platform/internal/application/commands"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/response"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
	spatial "github.com/samurenkoroma/agro-platform/internal/domain/spatial/repository"
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

	return h.uow.Execute(ctx, providers.NewSpatialProvider, func(provider repository.RepositoryProvider) (any, error) {
		spatialProvider, ok := provider.(spatial.SpatialProvider)
		if !ok {
			return nil, repository.ErrInvalidProviderType
		}

		unit := pu.New(
			vo.ID(orgId),
			cmd.ParentID,
			cmd.Type,
			cmd.Code,
			cmd.Name,
		)
		unit.Properties.AddCapabilities(cmd.Capabilities)
		if cmd.Dimensions != nil {
			unit.AddDimensions(cmd.Dimensions)
		}

		if err := spatialProvider.Units().Save(ctx, unit); err != nil {
			return nil, err
		}

		if cmd.CreateChild && cmd.Dimensions.CellCount != nil {
			for i := 1; i <= *cmd.Dimensions.CellCount; i++ {
				code := fmt.Sprintf("%s/S-%02d", cmd.Code, i)
				u := pu.New(
					vo.ID(orgId),
					&unit.ID,
					pu.Slot,
					code,
					&code,
				)
				if err := spatialProvider.Units().Save(ctx, u); err != nil {
					return nil, err
				}
				h.uow.RegisterAggregate(u)
			}
		}

		h.uow.RegisterAggregate(unit)
		return response.Id(unit.ID), nil
	})
}
