package mq

import (
	rmqclient "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"log/slog"
	"shortlink/internal/base"
	"time"
)

var (
	// maximum waiting time for receive func
	awaitDuration = time.Second * 5
	// maximum number of messages received at one time
	maxMessageNum int32 = 16
	// invisibleDuration should > 20s
	invisibleDuration = time.Second * 20
	// receive messages in a loop
)

func ConnectToRocketMQForProducer() (rmqclient.Producer, func()) {

	config := base.GetConfig().RocketMQ

	// In most case, you don't need to create many producers, singleton pattern is more recommended.
	producer, err := rmqclient.NewProducer(&rmqclient.Config{
		Endpoint: config.NameServer,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    config.AccessKey,
			AccessSecret: config.SecretKey,
		},
	},
		rmqclient.WithTopics(config.Topics...),
	)

	if err != nil {
		panic("Failed to create producer: " + err.Error())
	}
	if err = producer.Start(); err != nil {
		panic("Failed to start producer: " + err.Error())
	}

	// graceful stop producer
	stopFn := func() {
		if err := producer.GracefulStop(); err != nil {
			slog.Error("Failed to graceful stop producer", "error", err)
		}
	}

	return producer, stopFn
}

func ConnectToRocketMQForConsumer() (rmqclient.SimpleConsumer, func()) {

	config := base.GetConfig().RocketMQ

	subs := make(map[string]*rmqclient.FilterExpression)
	for _, topic := range config.Topics {
		subs[topic] = rmqclient.SUB_ALL
	}

	// In most case, you don't need to create many consumers, singleton pattern is more recommended.
	simpleConsumer, err := rmqclient.NewSimpleConsumer(&rmqclient.Config{
		Endpoint:      config.NameServer,
		NameSpace:     config.NameSpace,
		ConsumerGroup: config.ConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    config.AccessKey,
			AccessSecret: config.SecretKey,
		},
	},
		rmqclient.WithAwaitDuration(awaitDuration),
		rmqclient.WithSubscriptionExpressions(subs),
	)
	if err != nil {
		panic("Failed to create simple consumer: " + err.Error())
	}
	if err = simpleConsumer.Start(); err != nil {
		panic("Failed to start simple consumer: " + err.Error())
	}

	stopFn := func() {
		if err := simpleConsumer.GracefulStop(); err != nil {
			slog.Error("Failed to graceful stop simple consumer", "error", err)
		}
	}

	return simpleConsumer, stopFn
}
