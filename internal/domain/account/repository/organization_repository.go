package repository

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/account/aggregate/organization"
)

// OrganizationRepository интерфейс репозитория организаций
type OrganizationRepository interface {
	// Базовые операции
	Save(ctx context.Context, org *organization.Organization) error
	Update(ctx context.Context, org *organization.Organization) error
	FindByID(ctx context.Context, id string) (*organization.Organization, error)
	Delete(ctx context.Context, id string) error

	// Поиск
	FindByName(ctx context.Context, name string) (*organization.Organization, error)
	FindByTaxID(ctx context.Context, taxID string) (*organization.Organization, error)

	// Списки
	List(ctx context.Context, filter OrganizationFilter) ([]*organization.Organization, int, error)
	ListByUser(ctx context.Context, userID string) ([]*organization.Organization, error)

	// Статус
	Activate(ctx context.Context, id string) error
	Deactivate(ctx context.Context, id string) error
}

// OrganizationFilter фильтр для организаций
type OrganizationFilter struct {
	Search   string `json:"search,omitempty"`
	IsActive *bool  `json:"is_active,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	Offset   int    `json:"offset,omitempty"`
}
