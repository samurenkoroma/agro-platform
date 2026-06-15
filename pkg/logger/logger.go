package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
)

type contextKey struct{}

// New создаёт JSON-логгер. level: "debug" | "info" | "warn" | "error"
func New(level string, w io.Writer) *slog.Logger {
	if w == nil {
		w = os.Stdout
	}
	return slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level:     parseLevel(level),
		AddSource: true,
	}))
}

// WithContext кладёт логгер в контекст.
func WithContext(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, l)
}

// FromContext достаёт логгер из контекста.
// Если логгер не установлен — возвращает дефолтный (не паникует).
func FromContext(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(contextKey{}).(*slog.Logger); ok && l != nil {
		return l
	}
	return slog.Default()
}

func parseLevel(s string) slog.Level {
	switch s {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
