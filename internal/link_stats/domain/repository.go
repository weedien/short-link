package domain

import (
	"context"
	"shortlink/internal/link/domain/event"
)

type Repository interface {
	SaveLinkStats(ctx context.Context, statsInfo event.UserVisitInfo) error
}
