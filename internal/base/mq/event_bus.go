package mq

import (
	"context"
	"errors"
	rmqclient "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/bytedance/sonic"
	"log/slog"
	"reflect"
	"shortlink/internal/base/base_event"
	"time"
)

type RocketMqBasedEventBus struct {
	// topic+tag -> listener
	listenerMap  map[string][]base_event.EventListener
	typeRegistry map[string]reflect.Type
	producer     rmqclient.Producer
	consumer     rmqclient.SimpleConsumer
	stopFns      []func()
	mode         RunMode
}

type RunMode int

const (
	ConsumerMode RunMode = 1
	ProducerMode RunMode = 2
	MixMode      RunMode = 3
)

func NewRocketMqBasedEventBus(ctx context.Context, mode RunMode) *RocketMqBasedEventBus {

	var producer rmqclient.Producer
	var producerStopFn func()
	var consumer rmqclient.SimpleConsumer
	var consumerStopFn func()

	var stopFns []func()
	if mode == ProducerMode || mode == MixMode {
		producer, producerStopFn = ConnectToRocketMQForProducer()
		stopFns = append(stopFns, producerStopFn)
	}
	if mode == ConsumerMode || mode == MixMode {
		consumer, consumerStopFn = ConnectToRocketMQForConsumer()
		stopFns = append(stopFns, consumerStopFn)
	}

	bus := &RocketMqBasedEventBus{
		listenerMap:  make(map[string][]base_event.EventListener),
		typeRegistry: make(map[string]reflect.Type),
		producer:     producer,
		consumer:     consumer,
		stopFns:      stopFns,
	}

	if mode == ConsumerMode || mode == MixMode {
		go bus.startReceivingMessages(ctx)
	}

	return bus
}

func (bus *RocketMqBasedEventBus) Close() {
	for _, stopFn := range bus.stopFns {
		stopFn()
	}
}

func (bus *RocketMqBasedEventBus) startReceivingMessages(ctx context.Context) {
	for {
		mvs, err := bus.consumer.Receive(ctx, maxMessageNum, invisibleDuration)
		if err != nil {
			// todo 忽略消息为空时的异常
			slog.Error("Failed to receive message from RocketMQ", "error", err)
			continue
		}

		for _, mv := range mvs {

			idx := ""
			if mv.GetTag() != nil {
				idx = mv.GetTopic() + ":" + *mv.GetTag()
			} else {
				idx = mv.GetTopic()
			}

			var ok bool
			var listeners []base_event.EventListener
			if listeners, ok = bus.listenerMap[idx]; !ok {
				// 当前消息无人订阅
				continue
			}

			for _, listener := range listeners {
				go func(listener base_event.EventListener) {
					if err := listener.Process(ctx, string(mv.GetBody())); err != nil {
						slog.Error("Failed to process message", "error", err)
						return
					}
					// 消费成功 之后MQ将不会投递该消息 否则会进行重试
					if err := bus.consumer.Ack(ctx, mv); err != nil {
						slog.Error("Failed to ack message", "error", err)
						return
					}
				}(listener)
			}
		}
	}
}

func (bus *RocketMqBasedEventBus) Publish(ctx context.Context, event base_event.Event) error {

	if bus.mode == ConsumerMode {
		return errors.New("can't publish event in ConsumerMode")
	}

	marshal, err := sonic.Marshal(event)
	if err != nil {
		slog.Error("Failed to marshal event", "error", err)
		return err
	}

	msg := &rmqclient.Message{
		Topic: event.Topic(),
		Body:  marshal,
	}
	if event.Tag() != "" {
		msg.SetTag(event.Tag())
	}
	if event.Keys() != "" {
		msg.SetKeys(event.Keys())
	}
	if event.Delay() > 0 {
		msg.SetDelayTimestamp(time.Now().Add(event.Delay()))
	}
	if event.MessageGroup() != "" {
		// 用于保证顺序消费
		msg.SetMessageGroup(event.MessageGroup())
	}

	resp, err := bus.producer.Send(ctx, msg)
	if err != nil {
		slog.Error("Failed to send message to RocketMQ", "error", err, "msg", marshal)
		return err
	}
	slog.Info("Sent message to RocketMQ", "resp", resp)
	return nil
}

func (bus *RocketMqBasedEventBus) Subscribe(topic string, tag *string, listener base_event.EventListener) error {

	if bus.mode == ProducerMode {
		return errors.New("can't subscribe event in ProducerMode")
	}

	idx := ""
	if tag != nil {
		idx = topic + ":" + *tag
	} else {
		idx = topic
	}
	listeners, exists := bus.listenerMap[idx]
	if !exists {
		listeners = make([]base_event.EventListener, 0)
		bus.listenerMap[idx] = listeners
	}
	listeners = append(listeners, listener)
	return nil
}
