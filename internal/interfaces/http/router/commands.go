package router

import (
	"net/http"

	httpHandlers "github.com/samurenkoroma/agro-platform/internal/interfaces/http/handlers"
	infrastructure "github.com/samurenkoroma/agro-platform/internal/shared/di"
)

func RegisterCommands(mux *http.ServeMux, container infrastructure.Container) {
	handler := httpHandlers.NewCommandHandler(container.CommandBus)
	mux.Handle("/commands", handler)
}
