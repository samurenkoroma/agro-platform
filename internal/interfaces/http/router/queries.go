package router

import (
	"net/http"

	httpHandlers "github.com/samurenkoroma/agro-platform/internal/interfaces/http/handlers"
	infrastructure "github.com/samurenkoroma/agro-platform/internal/shared/di"
)

func RegisterQueries(mux *http.ServeMux, container infrastructure.Container) {
	handler := httpHandlers.NewQueryHandler(container.QueryBus)

	mux.Handle("/queries", handler)
}
