package crop

import (
	"context"
	"fmt"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	domain "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
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
	Key string `form:"key,omitempty"`
}

func (h *CropHandler) Ask(ctx context.Context, query any) (any, error) {
	return h.uow.Execute(ctx, providers.NewAgronomyProvider, func(provider repository.RepositoryProvider) (any, error) {
		agronomyProvider, ok := provider.(domain.AgronomyProvider)
		if !ok {
			return nil, fmt.Errorf("expected FarmProvider, got %T", provider)
		}
		return agronomyProvider.Crops().GetAll(ctx)
	})

}
