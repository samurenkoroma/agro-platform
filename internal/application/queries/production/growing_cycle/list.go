package growingcycle

import (
	"context"
	"errors"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type listHandler struct {
	cycles Projection
}

func NewList(cycles Projection) queries.Handler {
	return &listHandler{
		cycles: cycles,
	}
}

type ListQuery struct {
	Id *string `json:"id,omitempty"`
}

func (h *listHandler) Ask(ctx context.Context, payload any) (any, error) {
	cmd, ok := payload.(*ListQuery)
	if !ok {
		return nil, queries.ErrInvalidQueryType
	}
	orgID, ok := ctx.Value("organization_id").(string)
	if !ok {
		return nil, errors.New("organization_id is required")
	}
	return h.cycles.List(ctx, FilterCycle{
		OwnerId: vo.ID(orgID),
		UnitId:  cmd.Id,
	})
}
