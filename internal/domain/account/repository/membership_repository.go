package repository

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/account/aggregate/organization"
)

// MembershipRepository интерфейс репозитория членств
type MembershipRepository interface {
	// Базовые операции
	Save(ctx context.Context, membership *organization.Membership) error
	Update(ctx context.Context, membership *organization.Membership) error
	FindByID(ctx context.Context, id string) (*organization.Membership, error)
	Delete(ctx context.Context, id string) error

	// Поиск по связям
	FindByUser(ctx context.Context, userID string) ([]*organization.Membership, error)
	FindByOrganization(ctx context.Context, orgID string) ([]*organization.Membership, error)
	FindByUserAndOrganization(ctx context.Context, userID, orgID string) (*organization.Membership, error)

	// Проверка существования
	Exists(ctx context.Context, userID, orgID string) (bool, error)

	// Списки
	List(ctx context.Context, filter MembershipFilter) ([]*organization.Membership, int, error)

	// Управление статусом
	Activate(ctx context.Context, id string) error
	Deactivate(ctx context.Context, id string) error
	ChangeRole(ctx context.Context, id string, newRole organization.OrganizationRole) error
}

// MembershipFilter фильтр для членств
type MembershipFilter struct {
	UserID         string                        `json:"user_id,omitempty"`
	OrganizationID string                        `json:"organization_id,omitempty"`
	Role           organization.OrganizationRole `json:"role,omitempty"`
	IsActive       *bool                         `json:"is_active,omitempty"`
	Limit          int                           `json:"limit,omitempty"`
	Offset         int                           `json:"offset,omitempty"`
}
