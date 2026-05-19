package inmemory

import "context"

type Transaction struct {
	committed  bool
	rolledBack bool
}

func (t *Transaction) Commit(ctx context.Context) error {
	t.committed = true
	return nil
}

func (t *Transaction) Rollback(ctx context.Context) error {
	t.rolledBack = true

	return nil
}
