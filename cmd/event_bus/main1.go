package main

//
//import (
//	"context"
//	"log/slog"
//	"sync"
//	"time"
//)
//
//// AppEvent is an interface for events in the application layer.
//type AppEvent interface {
//	Tag() string
//	OccurredAt() time.Time
//}
//
//type CommonEvent struct {
//	occurredAt time.Time
//}
//
//func NewCommonEvent() CommonEvent {
//	return CommonEvent{
//		occurredAt: time.Now(),
//	}
//}
//
//func (e CommonEvent) OccurredAt() time.Time {
//	return e.occurredAt
//}
//
//// AppEventConsumer is an interface for consuming events in the application layer.
//type AppEventConsumer[E AppEvent] interface {
//	Process(ctx context.Context, e E) error
//}
//
//// AppEventBus is an interface for event bus in the application layer.
//type AppEventBus[E AppEvent] interface {
//	Publish(ctx context.Context, event E)
//	Subscribe(eventName string, consumer AppEventConsumer[E])
//}
//
//type SimpleEventBus[E AppEvent] struct {
//	consumers map[string][]AppEventConsumer[E]
//	mu        sync.RWMutex
//}
//
//func NewSimpleEventBus[E AppEvent]() *SimpleEventBus[E] {
//	return &SimpleEventBus[E]{
//		consumers: make(map[string][]AppEventConsumer[E]),
//	}
//}
//
//func (bus *SimpleEventBus[E]) Publish(ctx context.Context, event E) {
//	bus.mu.RLock()
//	defer bus.mu.RUnlock()
//
//	if consumers, found := bus.consumers[event.Tag()]; found {
//		for _, consumer := range consumers {
//			go func(consumer AppEventConsumer[E]) {
//				if err := consumer.Process(ctx, event); err != nil {
//					slog.Info("event consumer error", err)
//				}
//			}(consumer)
//		}
//	}
//}
//
//func (bus *SimpleEventBus[E]) Subscribe(eventName string, consumer AppEventConsumer[E]) {
//	bus.mu.Lock()
//	defer bus.mu.Unlock()
//	bus.consumers[eventName] = append(bus.consumers[eventName], consumer)
//}
//
//// -------------- LinkAccessedEvent --------------
//
//type LinkAccessedEvent struct {
//	CommonEvent
//	info string
//}
//
//func (e LinkAccessedEvent) Tag() string {
//	return "LinkAccessedEvent"
//}
//
//func NewLinkAccessedEvent(recordInfo string) LinkAccessedEvent {
//	return LinkAccessedEvent{
//		CommonEvent: NewCommonEvent(),
//		info:        recordInfo,
//	}
//}
//
//type LinkAccessedEventConsumer struct {
//}
//
//func (h LinkAccessedEventConsumer) Process(ctx context.Context, e LinkAccessedEvent) error {
//	slog.Info("LinkAccessedEventConsumer", e.info)
//	return nil
//}
//
//func main() {
//	bus := NewSimpleEventBus[AppEvent]()
//	bus.Subscribe("LinkAccessedEvent", LinkAccessedEventConsumer{})
//	bus.Publish(context.Background(), NewLinkAccessedEvent("test"))
//}
