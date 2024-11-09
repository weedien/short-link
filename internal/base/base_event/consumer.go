package base_event

import (
	"context"
)

// EventListener is an interface for consuming events in the application layer.
type EventListener interface {
	Process(ctx context.Context, e string) error
}
