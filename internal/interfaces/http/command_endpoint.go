package http

import (
	"encoding/json"
	"net/http"

	"github.com/samurenkoroma/agro-platform/internal/application/commands"
	"github.com/samurenkoroma/agro-platform/internal/interfaces/http/response"
	"github.com/samurenkoroma/agro-platform/pkg/logger"
)

type CommandPayload struct {
	Command string          `json:"command"`
	Data    json.RawMessage `json:"data"`
}

func CommandEndpoint(router commands.Router) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log := logger.FromContext(r.Context())
		var payload CommandPayload

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			log.Warn("command: invalid payload", "error", err)
			response.WriteValidationError(w, "invalid request payload: "+err.Error())
			return
		}
		if payload.Command == "" {
			log.Warn("command: empty command name")
			response.WriteValidationError(w, "command name is required")
			return
		}

		log = log.With("command", payload.Command)

		handlerCmd, err := router.ResolveCommandPayload(payload.Command, payload.Data)
		if err != nil {
			log.Warn("command: failed to resolve", "error", err)
			response.WriteError(w, http.StatusBadRequest, response.CodeBadRequest,
				"failed to decode command: "+err.Error())
			return
		}

		result, err := router.Dispatch(r.Context(), payload.Command, handlerCmd)
		if err != nil {
			log.Warn("command: dispatch failed", "error", err)
			// Используем стандартный ответ с ошибкой
			resp := response.FromError(err)
			resp.WriteJSON(w, getStatusCodeForError(resp.Error.Code))
			return
		}

		log.Info("command: executed")

		// Если команда вернула результат (например, LoginResponse)
		if result != nil {
			response.WriteSuccess(w, result)
		} else {
			// Успешное выполнение без данных
			response.WriteSuccess(w, map[string]string{"status": "ok"})
		}
	}
}

// getStatusCodeForError возвращает HTTP статус код по коду ошибки
func getStatusCodeForError(errorCode string) int {
	switch errorCode {
	case response.CodeBadRequest, response.CodeValidation:
		return http.StatusBadRequest
	case response.CodeUnauthorized:
		return http.StatusUnauthorized
	case response.CodeForbidden:
		return http.StatusForbidden
	case response.CodeNotFound:
		return http.StatusNotFound
	case response.CodeConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
