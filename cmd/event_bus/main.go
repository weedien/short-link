package main

import (
	"context"
	"log/slog"
	"reflect"
	"sync"
	"time"
)

// AppEvent is an interface for events in the application layer.
type AppEvent interface {
	Name() string
	OccurredAt() time.Time
}

type CommonEvent struct {
	occurredAt time.Time
}

func NewCommonEvent() CommonEvent {
	return CommonEvent{
		occurredAt: time.Now(),
	}
}

func (e CommonEvent) OccurredAt() time.Time {
	return e.occurredAt
}

// AppEventConsumer is an interface for consuming events in the application layer.
type AppEventConsumer interface {
	Process(ctx context.Context, e AppEvent) error
}

// AppEventBus is an interface for event bus in the application layer.
type AppEventBus interface {
	Publish(ctx context.Context, event AppEvent)
	Subscribe(eventType reflect.Type, consumer AppEventConsumer)
}

type SimpleEventBus struct {
	consumers map[reflect.Type][]AppEventConsumer
	mu        sync.RWMutex
}

func NewSimpleEventBus() AppEventBus {
	return &SimpleEventBus{
		consumers: make(map[reflect.Type][]AppEventConsumer),
	}
}

func (bus *SimpleEventBus) Publish(ctx context.Context, event AppEvent) {
	bus.mu.RLock()
	defer bus.mu.RUnlock()

	eventType := reflect.TypeOf(event)
	if consumers, found := bus.consumers[eventType]; found {
		for _, consumer := range consumers {
			go func(consumer AppEventConsumer) {
				if err := consumer.Process(ctx, event); err != nil {
					slog.Info("event consumer error", err)
				}
			}(consumer)
		}
	} else {
		slog.Warn("no consumer found for event", "event", event.Name())
	}
}

func (bus *SimpleEventBus) Subscribe(eventType reflect.Type, consumer AppEventConsumer) {
	bus.mu.Lock()
	defer bus.mu.Unlock()
	bus.consumers[eventType] = append(bus.consumers[eventType], consumer)
}

// -------------- LinkAccessedEvent --------------

type LinkAccessedEvent struct {
	CommonEvent
	info string
}

func (e LinkAccessedEvent) Name() string {
	return "LinkAccessedEvent"
}

func NewLinkAccessedEvent(recordInfo string) LinkAccessedEvent {
	return LinkAccessedEvent{
		CommonEvent: NewCommonEvent(),
		info:        recordInfo,
	}
}

type LinkAccessedEventConsumer struct {
}

func (h LinkAccessedEventConsumer) Process(ctx context.Context, e AppEvent) error {
	event := e.(LinkAccessedEvent)
	slog.Info("LinkAccessedEventConsumer", "info", event.info, "event", event)
	return nil
}

func main() {
	bus := NewSimpleEventBus()
	bus.Subscribe(reflect.TypeOf(LinkAccessedEvent{}), LinkAccessedEventConsumer{})
	bus.Publish(context.Background(), NewLinkAccessedEvent("test"))

	time.Sleep(time.Second)
}
