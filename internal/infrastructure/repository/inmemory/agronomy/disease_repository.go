package agronomy

import "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"

type diseaseRepository struct{}

func NewDiseaseRepository() repository.DiseaseRepository {
	return &diseaseRepository{}
}
