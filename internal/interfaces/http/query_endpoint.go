package http

import (
	"encoding/json"
	"net/http"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	"github.com/samurenkoroma/agro-platform/internal/interfaces/http/response"
)

type QueryPayload struct {
	Query string          `json:"query"`
	Data  json.RawMessage `json:"data"`
}

func QueryEndpoint(router queries.Router) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var payload QueryPayload

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "invalid request payload", http.StatusBadRequest)
			return
		}

		result, err := router.Dispatch(
			r.Context(),
			payload.Query,
			payload.Data,
		)
		if err != nil {
			resp := response.FromError(err)
			statusCode := getStatusCodeForError(resp.Error.Code)
			resp.WriteJSON(w, statusCode)
			return
		}

		response.WriteSuccess(w, result)
	}
}
