package bootstrap

import (
	msg "github.com/samurenkoroma/agro-platform/internal/infrastructure/messaging/inmemory"
	repoinfra "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory"
	txinfra "github.com/samurenkoroma/agro-platform/internal/infrastructure/tx/inmemory"
	infrastructure "github.com/samurenkoroma/agro-platform/internal/shared/di"
)

func NewContainer() *infrastructure.Container {
	registry := msg.NewRegistry()

	return &infrastructure.Container{
		CommandBus:   msg.NewCommandBus(registry),
		QueryBus:     msg.NewQueryBus(registry),
		EventBus:     msg.NewEventBus(),
		Repositories: repoinfra.NewProvider(),
		TxManager:    txinfra.NewManager(),
	}
}
