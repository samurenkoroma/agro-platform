package createproductionunit

import (
	"context"

	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"

	repository "github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type Handler struct {
	repositories repository.Provider
}

func NewHandler(repositories repository.Provider) *Handler {
	return &Handler{
		repositories: repositories,
	}
}

func (h *Handler) Handle(ctx context.Context, cmd Command) error {

	unit, err := pu.New(
		cmd.FarmID,
		cmd.Type,
		cmd.Name,
	)

	if cmd.ParentID != nil {
		unit.AssignParent(
			*cmd.ParentID,
		)
	}

	return h.
		repositories.
		Spatial().
		ProductionUnits().
		Save(
			unit,
		)
}
