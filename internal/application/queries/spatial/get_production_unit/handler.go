package getproductionunit

import (
	"context"

	repository "github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type Handler struct {
	repositories repository.Provider
}

func NewHandler(
	repositories repository.Provider,
) *Handler {

	return &Handler{
		repositories: repositories,
	}
}

func (
	h *Handler,
) Handle(
	ctx context.Context,

	query Query,
) (
	Result,
	error,
) {

	unit,
		err :=
		h.repositories.
			Spatial().
			ProductionUnits().
			GetByID(
				query.ID,
			)

	if err != nil {
		return Result{},
			err
	}

	if unit == nil {
		return Result{},
			ErrProductionUnitNotFound
	}

	result := Result{
		ID: unit.ID.String(),

		FarmID: unit.FarmID.String(),

		Name: unit.Name,

		Type: string(
			unit.Type,
		),

		CreatedAt: unit.CreatedAt.
			String(),

		UpdatedAt: unit.UpdatedAt.
			String(),
	}

	if unit.ParentID != nil {

		id :=
			unit.ParentID.
				String()

		result.ParentID =
			&id
	}

	return result,
		nil
}
