package getproductionunit

import (
	"context"
	"fmt"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
)

type QueryHandler struct {
}

func NewProductionUnitHandler() queries.Handler {
	return &QueryHandler{}
}

type GetCurrentFarmQuery struct {
	Id string `json:"id,omitempty"`
}

func (h *QueryHandler) Ask(ctx context.Context, payload any) (any, error) {
	q, ok := payload.(*GetCurrentFarmQuery)
	if !ok {
		return nil, queries.ErrInvalidPayloadType
	}

	return fmt.Sprintf("hello %s", q.Id), nil
}
