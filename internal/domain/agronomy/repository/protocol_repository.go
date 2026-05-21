package repository

import (
	"context"

	protocol "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop_protocol"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type CropProtocolRepository interface {
	Save(ctx context.Context, root *protocol.CropProtocol) error
	GetByID(ctx context.Context, id vo.ID) (*protocol.CropProtocol, error)
	GetByCrop(ctx context.Context, cropID vo.ID) ([]*protocol.CropProtocol, error)
	Exists(ctx context.Context, id vo.ID) (bool, error)
}
