package handlers

import (
	"net/http"

	qb "github.com/samurenkoroma/agro-platform/internal/shared/bus/query"
)

type QueryHandler struct {
	bus qb.Bus
}

func NewQueryHandler(bus qb.Bus) *QueryHandler {
	return &QueryHandler{
		bus: bus,
	}
}

func (h *QueryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
}
