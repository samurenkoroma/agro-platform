package queries

import (
	"context"
	"errors"
)

// DecoderFunc преобразует сырой payload (json, grpc, etc)
// в конкретный query struct
type DecoderFunc func([]byte) (any, error)
type Handler interface {
	Ask(ctx context.Context, payload any) (any, error)
}

//type HandlerFunc func(ctx context.Context, payload any) (any, error)

type Router interface {
	// Register регистрирует query
	Register(string, Handler, DecoderFunc)
	// Dispatch выполняет query
	Dispatch(ctx context.Context, name string, payload []byte) (any, error)
}

var ErrQueryNotFound = errors.New("query not registered")
