package repository

import "github.com/samurenkoroma/agro-platform/internal/domain/analytics/repository"

type yieldRecordRepository struct {
}

func NewYieldRecordRepository() repository.YieldRecordRepository {
	return &yieldRecordRepository{}
}
