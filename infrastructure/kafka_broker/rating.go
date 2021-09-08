package kafka_broker

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"kafka_topic_reader/domain/broker"
)

type ratingConsumer struct {
	consumer *kafka.Consumer
}

func (rc *ratingConsumer) GetRatingMessage() (broker.RatingMessage, error) {
	rawRatingMessage, err := rc.consumer.ReadMessage(-1)

	if err == nil {
		fmt.Printf("Received raw message from kafka %s: %s\n", rawRatingMessage.TopicPartition, string(rawRatingMessage.Value))

		ratingMessage, err := unmarshalRawRatingMessage(rawRatingMessage)
		return ratingMessage, err
	} else {
		fmt.Printf("Consumer error: %v (%v)\n", err, rawRatingMessage)
		return broker.RatingMessage{}, err
	}
}

func unmarshalRawRatingMessage(rawRatingMessage *kafka.Message) (broker.RatingMessage, error) {
	var ratingMessage broker.RatingMessage
	err := json.Unmarshal(rawRatingMessage.Value, &ratingMessage)

	if err == nil {
		fmt.Printf("Parse raw message from kafka %+v\n", ratingMessage)
		return ratingMessage, nil
	} else {
		fmt.Printf("Parsing error kafka messagge %+v", err)
		return broker.RatingMessage{}, err
	}
}

func NewRatingConsumer(consumer *kafka.Consumer) broker.RatingConsumer {
	return &ratingConsumer{consumer: consumer}
}
