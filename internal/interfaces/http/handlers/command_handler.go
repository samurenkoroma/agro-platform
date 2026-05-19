package handlers

import (
	"net/http"

	cb "github.com/samurenkoroma/agro-platform/internal/shared/bus/command"
)

type CommandHandler struct {
	bus cb.Bus
}

func NewCommandHandler(bus cb.Bus) *CommandHandler {
	return &CommandHandler{
		bus: bus,
	}
}

func (h *CommandHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
