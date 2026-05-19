package repository

import "github.com/samurenkoroma/agro-platform/internal/domain/automation/repository"

type actuatorRepository struct {
}

func NewActuatorRepository() repository.ActuatorRepository {
	return &actuatorRepository{}
}
