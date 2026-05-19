package router

import (
	"net/http"

	infrastructure "github.com/samurenkoroma/agro-platform/internal/shared/di"
)

func New(
	container infrastructure.Container,
) http.Handler {

	mux := http.NewServeMux()

	RegisterCommands(
		mux,
		container,
	)

	RegisterQueries(
		mux,
		container,
	)

	return mux
}
