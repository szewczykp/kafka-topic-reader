package kafka_broker

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"kafka_topic_reader/domain/broker"
	"kafka_topic_reader/domain/broker/vo"
)

type ratingConsumer struct {
	consumer *kafka.Consumer
}

func (rc *ratingConsumer) GetRatingMessage() (vo.RatingMessage, error) {
	rawRatingMessage, err := rc.consumer.ReadMessage(-1)

	if err == nil {
		fmt.Printf("Received raw message from kafka %s: %s\n", rawRatingMessage.TopicPartition, string(rawRatingMessage.Value))
		ratingMessage, err := vo.CreateRatingMessageFromJSON(rawRatingMessage.Value)

		return ratingMessage, err
	} else {
		fmt.Printf("Consumer error: %v (%v)\n", err, rawRatingMessage)
		return vo.RatingMessage{}, err
	}
}

func NewRatingConsumer(consumer *kafka.Consumer) broker.RatingConsumer {
	return &ratingConsumer{consumer: consumer}
}
