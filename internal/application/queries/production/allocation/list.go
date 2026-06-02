package allocation

import (
	"context"
	"errors"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type listHandler struct {
	allocations Projection
}

func NewList(allocations Projection) queries.Handler {
	return &listHandler{
		allocations: allocations,
	}
}

type ListQuery struct {
	Id string `json:"id,omitempty"`
}

func (h *listHandler) Ask(ctx context.Context, payload any) (any, error) {
	_, ok := payload.(*ListQuery)
	if !ok {
		return nil, queries.ErrInvalidQueryType
	}
	orgID, ok := ctx.Value("organization_id").(string)
	if !ok {
		return nil, errors.New("organization_id is required")
	}
	return h.allocations.List(ctx, vo.ID(orgID))
}
