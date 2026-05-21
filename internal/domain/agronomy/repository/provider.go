package repository

type AgronomyProvider interface {
	Crops() CropRepository
	Varieties() VarietyRepository
	Protocols() ProtocolRepository
	Diseases() DiseaseRepository
	StressProfiles() StressRepository
}
