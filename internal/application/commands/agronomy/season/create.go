package season

import (
	"context"
	"errors"
	"time"

	command "github.com/samurenkoroma/agro-platform/internal/application/commands"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/season"
	agronomy "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type CreateSeasonCmd struct {
	StartDate   string `json:"startDate,format:date" validate:"required"`
	EndDate     string `json:"endDate,format:date" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Status      string `json:"status" validate:"required"`
	Description string `json:"description"`
}

func (h *Handler) Create(ctx context.Context, cmd any) (any, error) {
	c, ok := cmd.(*CreateSeasonCmd)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}
	startDate, err := time.Parse(time.RFC3339, c.StartDate)
	if err != nil {
		return nil, err
	}
	endDate, err := time.Parse(time.RFC3339, c.EndDate)
	if err != nil {
		return nil, err
	}

	orgId, ok := ctx.Value("organization_id").(string)
	if !ok {
		return nil, errors.New("organization_id is required")
	}
	userId, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, errors.New("user_id is required")
	}

	return h.uow.Execute(ctx, providers.NewAgronomyProvider, func(provider repository.RepositoryProvider) (any, error) {
		agronomyProvider, ok := provider.(agronomy.AgronomyProvider)
		if !ok {
			return nil, repository.ErrInvalidProviderType
		}

		newSeason, err := season.New(c.Name, startDate, endDate, season.SeasonStatus(c.Status), vo.ID(userId), vo.ID(orgId))
		if err != nil {
			return nil, err
		}
		err = agronomyProvider.Seasons().Save(ctx, newSeason)
		if err != nil {
			return nil, err
		}

		h.uow.RegisterAggregate(newSeason)

		return nil, nil
	})
}
