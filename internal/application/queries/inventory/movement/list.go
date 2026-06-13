package movement

import (
	"context"
	"errors"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type listHandler struct{ proj Projection }

func NewList(proj Projection) queries.Handler { return &listHandler{proj: proj} }

type ListQuery struct {
	ItemID *string `json:"itemId,omitempty"`
}

func (h *listHandler) Ask(ctx context.Context, payload any) (any, error) {
	q, ok := payload.(*ListQuery)
	if !ok {
		return nil, queries.ErrInvalidQueryType
	}
	orgID, ok := ctx.Value("organization_id").(string)
	if !ok {
		return nil, errors.New("organization_id is required")
	}
	var itemID *vo.ID
	if q.ItemID != nil {
		id := vo.ID(*q.ItemID)
		itemID = &id
	}
	return h.proj.List(ctx, vo.ID(orgID), itemID)
}
