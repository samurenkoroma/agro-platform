package task

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type getOneHandler struct{ proj Projection }

func NewGetOne(proj Projection) queries.Handler { return &getOneHandler{proj: proj} }

type GetOneQuery struct {
	ID string `json:"id" validate:"required"`
}

func (h *getOneHandler) Ask(ctx context.Context, payload any) (any, error) {
	q, ok := payload.(*GetOneQuery)
	if !ok {
		return nil, queries.ErrInvalidQueryType
	}
	return h.proj.Get(ctx, vo.ID(q.ID))
}
