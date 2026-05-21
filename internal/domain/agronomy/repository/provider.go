package repository

import "github.com/samurenkoroma/agro-platform/internal/shared/repository"

type AgronomyProvider interface {
	repository.RepositoryProvider
	Crops() CropRepository
	Varieties() VarietyRepository
	Protocols() CropProtocolRepository
	Diseases() DiseaseRepository
	StressProfiles() StressRepository
}
