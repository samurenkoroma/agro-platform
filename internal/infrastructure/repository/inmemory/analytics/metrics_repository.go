package repository

import "github.com/samurenkoroma/agro-platform/internal/domain/analytics/repository"

type metricsRepository struct {
}

func NewMetricsRepository() repository.MetricsRepository {
	return &metricsRepository{}
}
