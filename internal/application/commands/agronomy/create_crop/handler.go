package createcrop

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/commands/response"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop"
	agronomy "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type Command struct {
	Name           string  `json:"name" validate:"required"`
	ScientificName *string `json:"scientific_name"`
	Category       string  `json:"category"`
}

type Handler struct {
	uow uow.UnitOfWork
}

func NewCreateCropHandler(uow uow.UnitOfWork) *Handler {
	return &Handler{uow: uow}
}

func (h *Handler) Handle(ctx context.Context, payload any) (any, error) {

	cmd := payload.(*Command)

	return h.uow.Execute(ctx, providers.NewAgronomyProvider, func(provider repository.RepositoryProvider) (any, error) {
		agronomyProvider, ok := provider.(agronomy.AgronomyProvider)
		if !ok {
			return nil, repository.ErrInvalidProviderType
		}
		root := crop.New(cmd.Name, crop.CropCategory(cmd.Category))

		root.ScientificName = cmd.ScientificName

		err := agronomyProvider.Crops().Save(ctx, root)

		if err != nil {
			return nil, err
		}

		h.uow.RegisterAggregate(root)

		return response.Id(root.ID), nil
	})
}
