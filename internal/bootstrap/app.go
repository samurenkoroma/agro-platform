package bootstrap

import (
	"net/http"

	httpRouter "github.com/samurenkoroma/agro-platform/internal/interfaces/http/router"
)

type App struct {
	server *http.Server
}

func New() (*App, error) {
	container := NewContainer()

	router := httpRouter.New(container)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return &App{server: server}, nil
}

func (a *App) Run() error {
	return a.server.ListenAndServe()
}
