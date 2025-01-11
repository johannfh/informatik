package utils

import (
	"context"
	"log/slog"
)

type loggerKey int

var lk loggerKey

func LoggerFromContext(ctx context.Context) *slog.Logger {
	return ctx.Value(lk).(*slog.Logger)
}

func ContextWithLogger(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, lk, l)
}
