package base_event

import (
	"github.com/google/uuid"
	"time"
)

type Event interface {
	// Id 事件唯一标识 用于保证幂等性
	Id() string

	Tag() string

	Keys() string

	// Topic 事件主题
	Topic() string

	Delay() time.Duration

	// MessageGroup 保证消息的顺序性
	MessageGroup() string

	OccurredAt() time.Time
}

type CommonEvent struct {
	id         string
	occurredAt time.Time
}

func NewCommonEvent() CommonEvent {
	return CommonEvent{
		id:         uuid.New().String(),
		occurredAt: time.Now(),
	}
}

func (e CommonEvent) OccurredAt() time.Time {
	return e.occurredAt
}

func (e CommonEvent) Id() string {
	return e.id
}

func (e CommonEvent) Name() string {
	return "CommonEvent"
}

func (e CommonEvent) Tag() string {
	return ""
}

func (e CommonEvent) Keys() string {
	return ""
}

func (e CommonEvent) Topic() string {
	panic("implement me")
}

func (e CommonEvent) Delay() time.Duration {
	return 0
}

func (e CommonEvent) MessageGroup() string {
	return ""
}
