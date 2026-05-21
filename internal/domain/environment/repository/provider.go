package repository

type EnvironmentProvider interface {
	Sensors() SensorRepository
}
