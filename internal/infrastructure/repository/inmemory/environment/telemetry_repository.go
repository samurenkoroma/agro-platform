package repository

import "github.com/samurenkoroma/agro-platform/internal/domain/environment/repository"

type telemetryRepository struct {
}

func NewTelemetryRepository() repository.TelemetryRepository {
	return &telemetryRepository{}
}
