package infrastructure

import (
	repoinfra "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory"
	txinfra "github.com/samurenkoroma/agro-platform/internal/infrastructure/tx/inmemory"
	"github.com/samurenkoroma/agro-platform/internal/shared/bus/command"
	"github.com/samurenkoroma/agro-platform/internal/shared/bus/event"
	"github.com/samurenkoroma/agro-platform/internal/shared/bus/query"
)

type Container struct {
	CommandBus   command.Bus
	QueryBus     query.Bus
	EventBus     event.Bus
	Repositories *repoinfra.Provider
	TxManager    *txinfra.Manager
}
