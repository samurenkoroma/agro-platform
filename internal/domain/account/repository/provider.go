package repository

type AccountProvider interface {
	Users() UserRepository
	Memberships() MembershipRepository
	Organizations() OrganizationRepository
	ProviderName() string
}
