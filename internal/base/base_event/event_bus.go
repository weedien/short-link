package base_event

import (
	"context"
)

// EventBus is an interface for event bus in the application layer.
type EventBus interface {
	Publish(ctx context.Context, event Event) error
	Subscribe(topic string, tag *string, listener EventListener) error
}
