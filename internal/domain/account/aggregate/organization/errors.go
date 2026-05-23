package organization

import "errors"

var (
	// User errors
	ErrMembershipNotFound   = errors.New("membership not found")
	ErrOrganizationNotFound = errors.New("organization not found")
	ErrInvalidEmail         = errors.New("invalid email")
	ErrInvalidPassword      = errors.New("invalid password")
	ErrPasswordTooShort     = errors.New("password too short (min 6 characters)")
	ErrEmailRequired        = errors.New("email is required")
	ErrUsernameRequired     = errors.New("username is required")
	ErrPasswordRequired     = errors.New("password is required")
	ErrUserBlocked          = errors.New("user is blocked")
)
