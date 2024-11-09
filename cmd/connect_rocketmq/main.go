package main

import (
	"context"
	"fmt"
	rmqclient "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	Topic         = "test_topic"
	Endpoint      = "193.112.178.249:8081"
	ConsumerGroup = "test_group"
	AccessKey     = ""
	SecretKey     = ""
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

func main() {
	_ = os.Setenv("mq.consoleAppender.enabled", "true")
	rmqclient.ResetLogger()
	// In most case, you don't need to create many producers, singleton pattern is more recommended.
	producer, err := rmqclient.NewProducer(&rmqclient.Config{
		Endpoint: Endpoint,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    AccessKey,
			AccessSecret: SecretKey,
		},
	},
		rmqclient.WithTopics(Topic),
	)
	if err != nil {
		log.Fatal(err)
	}
	// start producer
	err = producer.Start()
	if err != nil {
		log.Fatal("Failed to start producer: ", err)
	}
	// graceful stop producer
	defer func(producer rmqclient.Producer) {
		err := producer.GracefulStop()
		if err != nil {
			fmt.Printf("Failed to graceful stop producer: %s\n", err)
		}
	}(producer)

	go func() {
		for i := 0; i < 10; i++ {
			// new a message
			msg := &rmqclient.Message{
				Topic: Topic,
				Body:  []byte("this is a message : " + strconv.Itoa(i)),
			}
			// set keys and tag
			msg.SetKeys("a", "b")
			msg.SetTag("ab")
			// send message in sync
			resp, err := producer.Send(context.TODO(), msg)
			if err != nil {
				log.Fatal(err)
			}
			for i := 0; i < len(resp); i++ {
				fmt.Printf("%#v\n", resp[i])
			}
			// wait a moment
			time.Sleep(time.Second * 1)
		}
	}()

	// In most case, you don't need to create many consumers, singleton pattern is more recommended.
	simpleConsumer, err := rmqclient.NewSimpleConsumer(&rmqclient.Config{
		Endpoint:      Endpoint,
		ConsumerGroup: ConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    AccessKey,
			AccessSecret: SecretKey,
		},
	},
		rmqclient.WithAwaitDuration(awaitDuration),
		rmqclient.WithSubscriptionExpressions(map[string]*rmqclient.FilterExpression{
			Topic: rmqclient.SUB_ALL,
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	// start simpleConsumer
	err = simpleConsumer.Start()
	if err != nil {
		log.Fatal(err)
	}
	// graceful stop simpleConsumer
	defer func(simpleConsumer rmqclient.SimpleConsumer) {
		err := simpleConsumer.GracefulStop()
		if err != nil {
			fmt.Printf("Failed to graceful stop simple consumer: %s\n", err)
		}
	}(simpleConsumer)

	go func() {
		for {
			fmt.Println("start receive message")
			mvs, err := simpleConsumer.Receive(context.TODO(), maxMessageNum, invisibleDuration)
			if err != nil {
				fmt.Println(err)
			}
			// ack message
			for _, mv := range mvs {
				err := simpleConsumer.Ack(context.TODO(), mv)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(string(mv.GetBody()))
			}
			fmt.Println("wait a moment")
			fmt.Println()
			time.Sleep(time.Second * 3)
		}
	}()
	// run for a while
	time.Sleep(time.Minute)
}
