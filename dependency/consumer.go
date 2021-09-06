package dependency

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"kafka_topic_reader/config"
	"kafka_topic_reader/domain/broker"
	"kafka_topic_reader/infrastructure/kafka_broker"
)

func NewRatingConsumer(configuration config.KafkaConfiguration) broker.RatingConsumer  {

	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers" : configuration.BootstrapServer,
		"group.id": configuration.GroupId,
		"auto.offset.reset": "earliest",
	})


	if err != nil {
		panic(err)
	}

	kafkaConsumer.Subscribe(configuration.ConsumerTopicName, nil)

	return kafka_broker.NewRatingConsumer(kafkaConsumer)
}
