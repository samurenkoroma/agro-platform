package createproductionunit

import (
	"context"

	command "github.com/samurenkoroma/agro-platform/internal/application/commands"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/response"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
	spatial "github.com/samurenkoroma/agro-platform/internal/domain/spatial/repository"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type CreateProductionUnitHandler struct {
	uow uow.UnitOfWork
}

func NewProductionUnitHandler(uow uow.UnitOfWork) command.Handler {
	return &CreateProductionUnitHandler{uow: uow}
}

type Command struct {
	FarmID   vo.ID                 `json:"farmId" validate:"required"`
	Name     string                `json:"name" validate:"required"`
	Type     pu.ProductionUnitType `json:"type" validate:"required"`
	ParentID *vo.ID                `json:"parentId"`
}

func (h *CreateProductionUnitHandler) Handle(ctx context.Context, payload any) (any, error) {
	cmd, ok := payload.(*Command)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}

	return h.uow.Execute(ctx, uow.InMemoryDeps(uow.Spatial), func(provider repository.RepositoryProvider) (any, error) {
		spatialProvider, ok := provider.(spatial.SpatialProvider)
		if !ok {
			return nil, repository.ErrInvalidProviderType
		}

		unit, err := pu.New(
			cmd.FarmID,
			cmd.Type,
			cmd.Name,
		)
		if err != nil {
			return nil, err
		}

		if err := spatialProvider.Units().Save(ctx, unit); err != nil {
			return nil, err
		}

		h.uow.RegisterAggregate(unit)
		return response.Id(unit.ID), nil
	})
}
