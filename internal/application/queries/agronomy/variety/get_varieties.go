package catalog

import (
	"context"
	"errors"
	"fmt"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	domain "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	"github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type varietyHandler struct {
	uow uow.UnitOfWork
}

func NewVarietyHandler(uow uow.UnitOfWork) queries.Handler {
	return &varietyHandler{
		uow: uow,
	}
}

type VarietiesQuery struct {
	CropKey string `json:"cropId,omitempty"` // tomato, eggplant, cucumber
	Id      string `form:"id,omitempty"`
}

func (h *varietyHandler) Ask(ctx context.Context, query any) (any, error) {
	q, ok := query.(*VarietiesQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}
	return h.uow.Execute(ctx, providers.NewAgronomyProvider, func(provider repository.RepositoryProvider) (any, error) {
		agronomyProvider, ok := provider.(domain.AgronomyProvider)
		if !ok {
			return nil, fmt.Errorf("expected FarmProvider, got %T", provider)
		}
		return agronomyProvider.Varieties().GetByCrop(ctx, valueobject.ID(q.CropKey))
	})
}
