package config

import (
	"github.com/magiconair/properties"
	"log"
)

type KafkaConfiguration struct {
	ConsumerTopicName string `properties: "ConsumerTopicName"`
	BootstrapServer string `properties: "BootstrapServer"`
	GroupId string `properties: "GroupId"`
}

func ReadKafkaConfigurationFromFile() KafkaConfiguration {
	configuration := properties.MustLoadFile("${PWD}/resources/kafkaConfig.properties", properties.UTF8)

	var kafkaConfiguration KafkaConfiguration

	if err := configuration.Decode(&kafkaConfiguration); err != nil {
		log.Fatal(err)
	}

	return kafkaConfiguration
}