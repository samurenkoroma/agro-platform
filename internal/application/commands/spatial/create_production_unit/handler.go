package create_production_unit

import (
	"context"
	"errors"

	command "github.com/samurenkoroma/agro-platform/internal/application/commands"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/response"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
	spatial "github.com/samurenkoroma/agro-platform/internal/domain/spatial/repository"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type Handler struct {
	uow uow.UnitOfWork
}

func New(uow uow.UnitOfWork) *Handler {
	return &Handler{uow: uow}
}

type CreateCommand struct {
	Name     string                `json:"name" validate:"required"`
	Type     pu.ProductionUnitType `json:"type" validate:"required"`
	ParentID *vo.ID                `json:"parentId"`
}

func (h *Handler) Create(ctx context.Context, payload any) (any, error) {
	cmd, ok := payload.(*CreateCommand)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}
	organization_id, ok := ctx.Value("organization_id").(string)
	if !ok {
		return nil, errors.New("organization_id is required")
	}
	return h.uow.Execute(ctx, providers.NewSpatialProvider, func(provider repository.RepositoryProvider) (any, error) {
		spatialProvider, ok := provider.(spatial.SpatialProvider)
		if !ok {
			return nil, repository.ErrInvalidProviderType
		}

		unit, err := pu.New(
			vo.ID(organization_id),
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
