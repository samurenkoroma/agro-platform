package inmemory

import (
	"context"

	tx "github.com/samurenkoroma/agro-platform/internal/shared/tx"
)

type Manager struct{}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Begin(ctx context.Context) (tx.Transaction, error) {
	return &Transaction{}, nil
}
