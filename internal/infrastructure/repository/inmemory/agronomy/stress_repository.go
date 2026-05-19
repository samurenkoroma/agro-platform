package agronomy

import "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"

type stressRepository struct{}

func NewStressRepository() repository.StressRepository {
	return &stressRepository{}
}
