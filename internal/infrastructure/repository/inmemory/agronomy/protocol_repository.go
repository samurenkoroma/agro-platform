package agronomy

import (
	"context"

	protocol "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop_protocol"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type protocolRepository struct{}

func (p protocolRepository) Save(ctx context.Context, root *protocol.CropProtocol) error {
	//TODO implement me
	panic("implement me")
}

func (p protocolRepository) GetByID(ctx context.Context, id vo.ID) (*protocol.CropProtocol, error) {
	//TODO implement me
	panic("implement me")
}

func (p protocolRepository) GetByCrop(ctx context.Context, cropID vo.ID) ([]*protocol.CropProtocol, error) {
	//TODO implement me
	panic("implement me")
}

func (p protocolRepository) Exists(ctx context.Context, id vo.ID) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewProtocolRepository() repository.CropProtocolRepository {
	return &protocolRepository{}
}
