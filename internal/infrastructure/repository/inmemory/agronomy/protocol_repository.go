package agronomy

import "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"

type protocolRepository struct{}

func NewProtocolRepository() repository.ProtocolRepository {
	return &protocolRepository{}
}
