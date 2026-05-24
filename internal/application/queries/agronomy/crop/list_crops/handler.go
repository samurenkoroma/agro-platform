package listcrops

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	"github.com/samurenkoroma/agro-platform/internal/application/queries/agronomy/crop"
)

type handler struct {
	crops crop.Projection
}

type Query struct {
	Search   *string `json:"search"`
	Category *string `json:"category"`
	Archived *bool   `json:"archived"`
}

func New(crops crop.Projection) queries.Handler {
	return &handler{
		crops: crops,
	}
}

func (h *handler) Ask(ctx context.Context, query any) (any, error) {
	q, ok := query.(*Query)
	if !ok {
		return nil, queries.ErrInvalidQueryType
	}
	return h.crops.List(ctx,
		crop.ListFilter{
			Search:   q.Search,
			Category: q.Category,
			Archived: q.Archived,
		},
	)
}
