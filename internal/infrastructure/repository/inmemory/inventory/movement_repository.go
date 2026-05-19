package repository

import "github.com/samurenkoroma/agro-platform/internal/domain/inventory/repository"

type movementRepository struct{}

func NewMovementRepository() repository.MovementRepository {
	return &movementRepository{}
}
