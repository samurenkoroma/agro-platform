package tx

import "context"

type Manager interface {
	Begin(ctx context.Context) (UnitOfWork, error)
}
