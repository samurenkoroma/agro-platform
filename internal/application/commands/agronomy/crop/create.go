package crop

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/commands/response"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop"
	agronomy "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type CreateCropCommand struct {
	Name string `json:"name" validate:"required"`

	ScientificName string `json:"scientific_name"`
	Category       string `json:"category" required:"true"`
	Family         string `json:"family" required:"true"`
}

func (h *Handler) Create(ctx context.Context, payload any) (any, error) {
	cmd := payload.(*CreateCropCommand)

	return h.uow.Execute(ctx, providers.NewAgronomyProvider, func(provider repository.RepositoryProvider) (any, error) {
		agronomyProvider, ok := provider.(agronomy.AgronomyProvider)
		if !ok {
			return nil, repository.ErrInvalidProviderType
		}

		exist, _ := agronomyProvider.Crops().Exists(ctx, cmd.ScientificName)
		if exist {
			return nil, ErrCropAlreadyExist
		}

		root := crop.New(cmd.Name, crop.CropCategory(cmd.Category), cmd.Family, cmd.ScientificName)

		root.ScientificName = cmd.ScientificName

		if err := agronomyProvider.Crops().Save(ctx, root); err != nil {
			return nil, err
		}

		h.uow.RegisterAggregate(root)

		return response.Id(root.ID), nil
	})
}
