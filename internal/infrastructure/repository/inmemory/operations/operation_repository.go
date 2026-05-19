package repository

import "github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"

type operationRepository struct {
}

func NewOperationRepository() repository.OperationRepository {
	return &operationRepository{}
}
