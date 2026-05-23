package repository

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/account/aggregate/user"
)

// UserRepository интерфейс репозитория пользователей
type UserRepository interface {
	// Базовые операции
	Save(ctx context.Context, user *user.User) error
	Update(ctx context.Context, user *user.User) error
	FindByID(ctx context.Context, id string) (*user.User, error)
	Delete(ctx context.Context, id string) error

	// Поиск по уникальным полям
	FindByEmail(ctx context.Context, email string) (*user.User, error)
	FindByUsername(ctx context.Context, username string) (*user.User, error)

	// Списки
	List(ctx context.Context, filter UserFilter) ([]*user.User, int, error)

	// Статус
	UpdateLastLogin(ctx context.Context, userID string) error
	UpdateCurrentOrganization(ctx context.Context, userID, organizationID string) error
}

// UserFilter фильтр для пользователей
type UserFilter struct {
	Search string          `json:"search,omitempty"`
	Status user.UserStatus `json:"status,omitempty"`
	Role   user.Role       `json:"role,omitempty"`
	OrgID  string          `json:"org_id,omitempty"`
	Limit  int             `json:"limit,omitempty"`
	Offset int             `json:"offset,omitempty"`
}
