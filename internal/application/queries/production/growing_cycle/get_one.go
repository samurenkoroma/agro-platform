package growingcycle

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
)

type handler struct {
	cycles Projection
}

func NewGetOne(cycles Projection) queries.Handler {
	return &handler{
		cycles: cycles,
	}
}

type GetOneQuery struct {
	Id string `json:"id" validate:"required"`
}

func (h *handler) Ask(ctx context.Context, payload any) (any, error) {
	_, ok := payload.(*GetOneQuery)
	if !ok {
		return nil, queries.ErrInvalidQueryType
	}

	return nil, nil //h.cycles.Get(ctx, vo.ID(q.Id))
}
