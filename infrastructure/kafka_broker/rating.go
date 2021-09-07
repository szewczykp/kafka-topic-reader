package kafka_broker

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"kafka_topic_reader/domain/broker"
)

type ratingConsumer struct {
	consumer *kafka.Consumer
}

func (rc *ratingConsumer) GetRatingMessage() {
	fmt.Println("Start receiving from Kafka")

	for {
		message, err := rc.consumer.ReadMessage(-1)

		if err == nil {
			fmt.Printf("Received from Kafka %s: %s\n", message.TopicPartition, string(message.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, message)
			break
		}
	}
}

func NewRatingConsumer(consumer *kafka.Consumer) broker.RatingConsumer {
	return &ratingConsumer{consumer: consumer}
}
