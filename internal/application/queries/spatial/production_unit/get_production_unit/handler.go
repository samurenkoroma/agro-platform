package getproductionunit

import (
	"context"
	"fmt"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
)

type QueryHandler struct {
}

func New(uow.DB) queries.Handler {
	return &QueryHandler{}
}

type Query struct {
	Id string `json:"id,omitempty"`
}

func (h *QueryHandler) Ask(ctx context.Context, payload any) (any, error) {
	q, ok := payload.(*Query)
	if !ok {
		return nil, queries.ErrInvalidQueryType
	}

	return fmt.Sprintf("hello %s", q.Id), nil
}
