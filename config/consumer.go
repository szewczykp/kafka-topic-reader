package config

import "os"

func GetConsumerConfig() (string, string, string, string) {
	var host, port, consumerTopicName, groupId string

	if host = os.Getenv("KAFKA_ADVERTISED_HOST_NAME"); len(host) == 0 {
		host = "localhost"
	}
	if port = os.Getenv("KAFKA_ADVERTISED_PORT"); len(port) == 0 {
		port = "9092"
	}
	if consumerTopicName = os.Getenv("KAFKA_CONSUMER_TOPIC_NAME"); len(consumerTopicName) == 0 {
		consumerTopicName = "test-topic"
	}
	if groupId = os.Getenv("KAFKA_CONSUMER_GROUP_ID"); len(groupId) == 0 {
		groupId = "foo"
	}

	return host, port, consumerTopicName, groupId
}
