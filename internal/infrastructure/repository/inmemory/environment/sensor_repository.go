package repository

import "github.com/samurenkoroma/agro-platform/internal/domain/environment/repository"

type sensorRepository struct {
}

func NewSensorRepository() repository.SensorRepository {
	return &sensorRepository{}
}
