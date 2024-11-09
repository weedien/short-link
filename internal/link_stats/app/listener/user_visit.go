package listener

import (
	"context"
	"shortlink/internal/base/base_event"
	"shortlink/internal/link/domain/event"
	"shortlink/internal/link_stats/domain"
)

type RecordLinkVisitListener struct {
	repo domain.Repository
}

func NewRecordLinkVisitListener(repo domain.Repository) RecordLinkVisitListener {
	return RecordLinkVisitListener{repo: repo}
}

func (h RecordLinkVisitListener) Process(ctx context.Context, e base_event.Event) error {
	if ve, ok := e.(event.UserVisitEvent); ok {
		return h.repo.SaveLinkStats(ctx, ve.VisitInfo)
	}
	return nil
}
