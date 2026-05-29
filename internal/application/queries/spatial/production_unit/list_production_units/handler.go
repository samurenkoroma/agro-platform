package listproductionunits

import (
	"context"
	"fmt"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	productionunit "github.com/samurenkoroma/agro-platform/internal/application/queries/spatial/production_unit"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	postgres "github.com/samurenkoroma/agro-platform/internal/infrastructure/projection/postgres/spatial/production_unit"
)

type handler struct {
	units productionunit.Projection
}

func New(db uow.DB) queries.Handler {
	return &handler{
		units: postgres.New(db),
	}
}

type Query struct {
	Id string `json:"id,omitempty"`
}

func (h *handler) Ask(ctx context.Context, payload any) (any, error) {
	q, ok := payload.(*Query)
	if !ok {
		return nil, queries.ErrInvalidQueryType
	}

	return fmt.Sprintf("hello %s", q.Id), nil
}
