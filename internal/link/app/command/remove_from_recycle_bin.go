package command

import (
	"context"
	"log/slog"
	"shortlink/internal/base/decorator"
	"shortlink/internal/base/metrics"
	"shortlink/internal/link/domain"
	"shortlink/internal/link/domain/link"
)

type removeFromRecycleBinHandler struct {
	repo domain.Repository
}

type RemoveFromRecycleBinHandler decorator.CommandHandler[link.Identifier]

func NewRemoveFromRecycleBinHandler(
	repo domain.Repository,
	logger *slog.Logger,
	metricsClient metrics.Client,
) RemoveFromRecycleBinHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators[link.Identifier](
		removeFromRecycleBinHandler{repo},
		logger,
		metricsClient,
	)
}

func (h removeFromRecycleBinHandler) Handle(ctx context.Context, id link.Identifier) error {
	return h.repo.RemoveFromRecycleBin(ctx, id)
}
