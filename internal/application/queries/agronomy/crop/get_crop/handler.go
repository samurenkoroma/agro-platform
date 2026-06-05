package getcrop

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	"github.com/samurenkoroma/agro-platform/internal/application/queries/agronomy/crop"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	projection "github.com/samurenkoroma/agro-platform/internal/infrastructure/projection/postgres/agronomy/crop"
)

type handler struct{ crops crop.Projection }
type Query struct {
	ID string `json:"id" validate:"required"`
}

func New(db uow.DB) queries.Handler {
	return &handler{crops: projection.New(db)}
}

func (h *handler) Ask(ctx context.Context, query any) (any, error) {
	q, ok := query.(*Query)
	if !ok {
		return nil, queries.ErrInvalidQueryType
	}
	return h.crops.Get(ctx, q.ID)
}
