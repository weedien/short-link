package decorator

import (
	"context"
	"fmt"
	"log/slog"
	"shortlink/internal/base/metrics"
	"strings"
)

func ApplyCommandDecorators[H any](handler CommandHandler[H], logger *slog.Logger, metrics metrics.Client) CommandHandler[H] {
	return commandLoggingDecorator[H]{
		base:   commandMetricsDecorator[H]{base: handler, client: metrics},
		logger: logger,
	}
}

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}

// CommandWithResult 我希望通过这个接口知道哪些命令是有返回值的
type CommandWithResult[R any] interface {
	ExecutionResult() R
}
