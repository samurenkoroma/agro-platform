package growingcycle

import (
	"context"
	"errors"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type summaryHandler struct {
	cycles Projection
}

func NewSummaryHandler(cycles Projection) queries.Handler {
	return &summaryHandler{
		cycles: cycles,
	}
}

type SummaryQuery struct {
	Id vo.ID `json:"id,omitempty"`
}

func (h *summaryHandler) Ask(ctx context.Context, payload any) (any, error) {
	cmd, ok := payload.(*SummaryQuery)
	if !ok {
		return nil, queries.ErrInvalidQueryType
	}
	orgID, ok := ctx.Value("organization_id").(string)
	if !ok {
		return nil, errors.New("organization_id is required")
	}
	return h.cycles.Summary(ctx, vo.ID(orgID), cmd.Id)
}
