package main

import (
	"github.com/sillyhatxu/mini-mq/client/client"
	"github.com/sillyhatxu/mini-mq/client/consumer"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	Address      = "localhost:8082"
	TopicName    = "test_topic"
	TopicGroup   = "test-5"
	Offset       = 0
	ConsumeCount = 5
)

var Client = client.NewClient("localhost:8200", client.Timeout(30*time.Second))

type ConsumerTest struct{}

func (ct ConsumerTest) MessageDelivery(delivery consumer.Delivery) error {
	logrus.Infof("delivery { TopicName:%v, TopicGroup:%v, LatestOffset:%v , data length : %d}", delivery.TopicName, delivery.TopicGroup, delivery.LatestOffset, len(delivery.TopicDataArray))

	return nil
}

func main() {
	consume := consumer.NewConsumerClient(Client, TopicName, TopicGroup, Offset, ConsumeCount)
	err := consume.Consume(&ConsumerTest{})
	if err != nil {
		panic(err)
	}
}
