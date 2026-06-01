package productionunit

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type handler struct {
	units Projection
}

func NewGetOne(units Projection) queries.Handler {
	return &handler{
		units: units,
	}
}

type GetOneQuery struct {
	Id string `json:"id" validate:"required"`
}

func (h *handler) Ask(ctx context.Context, payload any) (any, error) {
	q, ok := payload.(*GetOneQuery)
	if !ok {
		return nil, queries.ErrInvalidQueryType
	}

	return h.units.Get(ctx, vo.ID(q.Id))
}
