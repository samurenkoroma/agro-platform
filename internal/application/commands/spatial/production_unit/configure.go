package productionunit

import (
	"context"
	"errors"

	command "github.com/samurenkoroma/agro-platform/internal/application/commands"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/response"
	spatial2 "github.com/samurenkoroma/agro-platform/internal/application/services/spatial"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	spatial "github.com/samurenkoroma/agro-platform/internal/domain/spatial/repository"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

func (h *Handler) Configure(ctx context.Context, payload any) (any, error) {
	cmd, ok := payload.(*ConfigureCommand)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}

	orgId, ok := ctx.Value("organization_id").(string)
	if !ok {
		return nil, errors.New("organization_id is required")
	}

	return h.uow.Execute(ctx, providers.NewSpatialProvider, func(provider repository.RepositoryProvider, exec uow.Execution) (any, error) {
		spatialProvider, ok := provider.(spatial.SpatialProvider)
		if !ok {
			return nil, repository.ErrInvalidProviderType
		}

		parent, err := spatialProvider.Units().GetByID(ctx, cmd.Id, vo.ID(orgId))
		if err != nil {
			return nil, err
		}

		if err := spatial2.NewUnitLayoutGenerator(spatialProvider.Units(), exec).Generate(ctx, parent, cmd.Schema); err != nil {
			return nil, err
		}
		return response.Id(parent.ID), nil
	})
}
