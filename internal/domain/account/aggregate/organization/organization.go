package organization

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/samurenkoroma/agro-platform/internal/domain/account/aggregate/user"
	"github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
)

// Organization организация (хозяйство, ферма, предприятие)
type Organization struct {
	aggregate.BaseAggregate
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	TaxID     string    `json:"tax_id"` // ИНН
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewSimpleOrganization(name string) (*Organization, error) {
	return NewOrganization(name, "", "", "", "")
}

// NewOrganization создает новую организацию
func NewOrganization(name, taxID, address, phone, email string) (*Organization, error) {
	if name == "" {
		return nil, errors.New("organization name is required")
	}

	now := time.Now()
	return &Organization{
		ID:        uuid.New().String(),
		Name:      name,
		TaxID:     taxID,
		Address:   address,
		Phone:     phone,
		Email:     email,
		IsActive:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// OrganizationRole роль пользователя в организации
type OrganizationRole string

const (
	OrgRoleOwner   OrganizationRole = "owner"   // владелец (полный доступ)
	OrgRoleAdmin   OrganizationRole = "admin"   // администратор организации
	OrgRoleAgronom OrganizationRole = "agronom" // агроном
	OrgRoleWorker  OrganizationRole = "worker"  // рабочий
	OrgRoleViewer  OrganizationRole = "viewer"  // только просмотр
)

// OrganizationPermissions права в рамках организации
var OrganizationPermissions = map[OrganizationRole][]user.Permission{
	OrgRoleOwner: {
		user.PermPlanCreate, user.PermPlanRead, user.PermPlanUpdate, user.PermPlanDelete, user.PermPlanComplete,
		user.PermTaskCreate, user.PermTaskRead, user.PermTaskUpdate, user.PermTaskComplete,
		user.PermUserCreate, user.PermUserRead, user.PermUserUpdate, user.PermUserDelete,
		user.PermVarietyCreate, user.PermVarietyRead, user.PermVarietyUpdate,
		user.PermReportView, user.PermReportExport,
		"org:manage", "org:delete",
	},
	OrgRoleAdmin: {
		user.PermPlanCreate, user.PermPlanRead, user.PermPlanUpdate, user.PermPlanComplete,
		user.PermTaskCreate, user.PermTaskRead, user.PermTaskUpdate,
		user.PermUserCreate, user.PermUserRead, user.PermUserUpdate,
		user.PermVarietyRead,
		user.PermReportView, user.PermReportExport,
		"org:manage",
	},
	OrgRoleAgronom: {
		user.PermPlanCreate, user.PermPlanRead, user.PermPlanUpdate, user.PermPlanComplete,
		user.PermTaskCreate, user.PermTaskRead, user.PermTaskUpdate,
		user.PermVarietyRead,
		user.PermReportView, user.PermReportExport,
	},
	OrgRoleWorker: {
		user.PermPlanRead,
		user.PermTaskRead, user.PermTaskComplete,
		user.PermReportView,
	},
	OrgRoleViewer: {
		user.PermPlanRead,
		user.PermTaskRead,
		user.PermReportView,
	},
}

// HasPermission проверяет наличие права в организации
func (r OrganizationRole) HasPermission(perm user.Permission) bool {
	perms, ok := OrganizationPermissions[r]
	if !ok {
		return false
	}
	for _, p := range perms {
		if p == perm {
			return true
		}
	}
	return false
}
