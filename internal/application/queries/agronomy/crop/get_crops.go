package crop

import (
	"context"
	"fmt"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	domain "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type CropHandler struct {
	uow uow.UnitOfWork
}

func NewCropHandler(uow uow.UnitOfWork) queries.Handler {
	return &CropHandler{
		uow: uow,
	}
}

type CropsQuery struct {
	ID       string `json:"id,omitempty"`
	Search   string `json:"search,omitempty"`
	Category string `json:"category,omitempty"`
	Family   string `json:"family,omitempty"`
	Archived *bool  `json:"archived,omitempty"`
}

func (h *CropHandler) Ask(ctx context.Context, query any) (any, error) {
	q, ok := query.(*CropsQuery)
	if !ok {
		return nil, queries.ErrInvalidQueryType
	}

	return h.uow.Execute(ctx, providers.NewAgronomyProvider, func(provider repository.RepositoryProvider) (any, error) {
		agronomyProvider, ok := provider.(domain.AgronomyProvider)
		if !ok {
			return nil, fmt.Errorf("expected FarmProvider, got %T", provider)
		}
		if q.ID != "" {
			return agronomyProvider.Crops().GetByID(ctx, vo.ID(q.ID))
		}

		return agronomyProvider.Crops().GetAll(ctx)
	})

}
