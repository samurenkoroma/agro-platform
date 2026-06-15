package http

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/samurenkoroma/agro-platform/pkg/logger"
)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (sw *statusWriter) WriteHeader(code int) {
	sw.status = code
	sw.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler, base *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		requestID := uuid.New().String()

		log := base.With("request_id", requestID, "method", r.Method, "path", r.URL.Path)

		w.Header().Set("X-Request-ID", requestID)

		ctx := logger.WithContext(r.Context(), log)
		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(sw, r.WithContext(ctx))

		log.Info("http request", "status", sw.status, "duration_ms", time.Since(start).Milliseconds())
	})
}
