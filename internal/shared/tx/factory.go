package tx

import (
	"context"
	"database/sql"

	"github.com/samurenkoroma/agro-platform/internal/shared/bus"
)

type Factory interface {
	Begin(ctx context.Context) (UnitOfWork, error)
	DB() *sql.DB
}

type UoWFactory struct {
	db  *sql.DB
	bus bus.EventBus
}

func NewUnitOfWorkFactory(db *sql.DB, bus bus.EventBus) Factory {
	return &UoWFactory{
		db:  db,
		bus: bus,
	}
}

func (f *UoWFactory) DB() *sql.DB {
	return f.db
}

func (f *UoWFactory) Begin(ctx context.Context) (UnitOfWork, error) {
	tx, err := f.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return NewUnitOfWork(ctx, tx, f, f.bus), nil
}
