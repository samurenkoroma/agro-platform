package repository

import "errors"

var (
	ErrProviderNotFound   = errors.New("repository provider not found")
	ErrRepositoryNotFound = errors.New("repository not found")
)
