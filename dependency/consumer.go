package dependency

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"kafka_topic_reader/config"
	"kafka_topic_reader/domain/broker"
	"kafka_topic_reader/infrastructure/kafka_broker"
)

func NewRatingConsumer() broker.RatingConsumer {
	host, port, consumerTopicName, groupId := config.GetConsumerConfig()

	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", host, port),
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	subscriberError := kafkaConsumer.Subscribe(consumerTopicName, nil)

	if subscriberError != nil {
		panic(subscriberError)
	}

	return kafka_broker.NewRatingConsumer(kafkaConsumer)
}
