package timeline

import (
	"context"
	"errors"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type handler struct{ proj Projection }

func NewGet(proj Projection) queries.Handler { return &handler{proj: proj} }

type GetQuery struct {
	GrowingCycleID *string `json:"growingCycleId,omitempty"`
}

func (h *handler) Ask(ctx context.Context, payload any) (any, error) {
	q, ok := payload.(*GetQuery)
	if !ok {
		return nil, queries.ErrInvalidQueryType
	}
	orgID, ok := ctx.Value("organization_id").(string)
	if !ok {
		return nil, errors.New("organization_id is required")
	}
	var cycleID *vo.ID
	if q.GrowingCycleID != nil {
		id := vo.ID(*q.GrowingCycleID)
		cycleID = &id
	}
	return h.proj.Get(ctx, vo.ID(orgID), cycleID)
}
