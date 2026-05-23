package providers

import (
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	domain "github.com/samurenkoroma/agro-platform/internal/domain/account/repository"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/postgres/account"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
)

type accountProvider struct {
	db uow.DB

	users       domain.UserRepository
	orgs        domain.OrganizationRepository
	memberships domain.MembershipRepository
}

func (p *accountProvider) Users() domain.UserRepository {
	if p.users == nil {
		p.users = postgres.NewUserRepository(p.db)
	}
	return p.users
}

func (p *accountProvider) Memberships() domain.MembershipRepository {
	if p.memberships == nil {
		p.memberships = postgres.NewMembershipRepository(p.db)
	}
	return p.memberships
}

func (p *accountProvider) Organizations() domain.OrganizationRepository {
	if p.orgs == nil {
		p.orgs = postgres.NewOrganizationRepository(p.db)
	}
	return p.orgs
}

func (p *accountProvider) ProviderName() string {
	return "account"
}

func NewAccountProvider(db uow.DB) repository.RepositoryProvider {
	return &accountProvider{
		db: db,
	}
}
