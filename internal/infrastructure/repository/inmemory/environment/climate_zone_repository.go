package repository

import "github.com/samurenkoroma/agro-platform/internal/domain/environment/repository"

type climateZoneRepository struct {
}

func NewClimateZoneRepository() repository.ClimateZoneRepository {
	return &climateZoneRepository{}
}
