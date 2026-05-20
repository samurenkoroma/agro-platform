package createproductionunit

import (
	"context"
	"fmt"

	command "github.com/samurenkoroma/agro-platform/internal/application/commands"
	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/spatial"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
	"github.com/samurenkoroma/agro-platform/internal/shared/tx"
)

type CreateProductionUnitHandler struct {
	uowFactory tx.Factory
}

func NewProductionUnitHandler(uowFactory tx.Factory) command.Handler {
	return &CreateProductionUnitHandler{uowFactory: uowFactory}
}

func (h *CreateProductionUnitHandler) Handle(ctx context.Context, payload any) (any, error) {
	cmd, ok := payload.(*Command)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return uow.Execute(ctx, spatial.NewInMemorySpatialProvider, func(provider repository.RepositoryProvider) (any, error) {
		spatialProvider, ok := provider.(*spatial.InMemorySpatialProvider)
		if !ok {
			return nil, fmt.Errorf("expected InMemorySpatialProvider, got %T", provider)
		}

		unit, err := pu.New(
			cmd.FarmID,
			cmd.Type,
			cmd.Name,
		)
		if err != nil {
			return nil, err
		}

		spatialProvider.Units().Save(unit)

		uow.RegisterAggregate(unit)
		return fmt.Sprintf("created %v", unit.CreatedAt), nil
	})
}
