package http

import (
	"encoding/json"
	"net/http"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	"github.com/samurenkoroma/agro-platform/internal/interfaces/http/response"
	"github.com/samurenkoroma/agro-platform/pkg/logger"
)

type QueryPayload struct {
	Query string          `json:"query"`
	Data  json.RawMessage `json:"data"`
}

func QueryEndpoint(router queries.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := logger.FromContext(r.Context())

		var payload QueryPayload

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			log.Warn("query: invalid payload", "error", err)
			http.Error(w, "invalid request payload", http.StatusBadRequest)
			return
		}

		if payload.Query == "" {
			log.Warn("query: empty query name")
			response.WriteValidationError(w, "query name is required")
			return
		}

		log = log.With("query", payload.Query)

		result, err := router.Dispatch(r.Context(), payload.Query, payload.Data)

		if err != nil {
			log.Warn("query: dispatch failed", "error", err)
			resp := response.FromError(err)
			resp.WriteJSON(w, getStatusCodeForError(resp.Error.Code))
			return
		}

		log.Info("query: executed")
		response.WriteSuccess(w, result)
	}
}
