package command

import (
	"context"
	"log/slog"
	"shortlink/internal/base/decorator"
	"shortlink/internal/base/metrics"
	"shortlink/internal/link/domain"
	"shortlink/internal/link/domain/link"
)

type saveToRecycleBinHandler struct {
	repo domain.Repository
}

type SaveToRecycleBinHandler decorator.CommandHandler[link.Identifier]

func NewSaveToRecycleBinHandler(
	repo domain.Repository,
	logger *slog.Logger,
	metricsClient metrics.Client,
) SaveToRecycleBinHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators[link.Identifier](
		saveToRecycleBinHandler{repo},
		logger,
		metricsClient,
	)
}

func (h saveToRecycleBinHandler) Handle(ctx context.Context, id link.Identifier) error {
	return h.repo.SaveToRecycleBin(ctx, id)
}
