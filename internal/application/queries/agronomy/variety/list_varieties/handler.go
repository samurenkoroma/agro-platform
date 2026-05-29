package listvarieties

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	"github.com/samurenkoroma/agro-platform/internal/application/queries/agronomy/variety"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	projection "github.com/samurenkoroma/agro-platform/internal/infrastructure/projection/postgres/agronomy/variety"
)

type varietyHandler struct {
	varieties variety.Projection
}

func New(db uow.DB) queries.Handler {
	return &varietyHandler{
		varieties: projection.New(db),
	}
}

type Query struct {
	CropKey string `json:"cropKey,omitempty"` // tomato, eggplant, cucumber
}

func (h *varietyHandler) Ask(ctx context.Context, query any) (any, error) {
	q, ok := query.(*Query)
	if !ok {
		return nil, queries.ErrInvalidQueryType
	}

	return h.varieties.List(ctx, variety.ListFilter{
		CropKey: q.CropKey,
	})

}
