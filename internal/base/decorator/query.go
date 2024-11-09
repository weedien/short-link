package decorator

import (
	"context"
	"log/slog"
	"shortlink/internal/base/metrics"
)

func ApplyQueryDecorators[H any, R any](handler QueryHandler[H, R], logger *slog.Logger, metrics metrics.Client) QueryHandler[H, R] {
	return queryLoggingDecorator[H, R]{
		base: queryMetricsDecorator[H, R]{
			base:   handler,
			client: metrics,
		},
		logger: logger,
	}
}

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, q Q) (R, error)
}
