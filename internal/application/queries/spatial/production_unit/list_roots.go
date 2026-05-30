package productionunit

import (
	"context"
	"errors"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type listRootsHandler struct {
	units Projection
}

func NewListRoots(units Projection) queries.Handler {
	return &listRootsHandler{
		units: units,
	}
}

type Query struct {
	Id string `json:"id,omitempty"`
}

func (h *listRootsHandler) Ask(ctx context.Context, payload any) (any, error) {
	_, ok := payload.(*Query)
	if !ok {
		return nil, queries.ErrInvalidQueryType
	}
	orgID, ok := ctx.Value("organization_id").(string)
	if !ok {
		return nil, errors.New("organization_id is required")
	}
	return h.units.ListRoots(ctx, vo.ID(orgID))
}
