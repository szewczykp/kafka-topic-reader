package main

import (
	"fmt"
	"kafka_topic_reader/config"
	"kafka_topic_reader/dependency"
	"kafka_topic_reader/domain/interactor"
)

func main() {
	kafkaProperties := config.ReadKafkaConfigurationFromFile()
	db, err := dependency.NewPostgresConnection()

	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	} else {
		fmt.Println("Database connection was initialized.")
		defer dependency.Close(db)
	}

	ratingInteractor := interactor.NewRatingInteractor(
		dependency.NewRatingRepository(db),
		dependency.NewRatingConsumer(kafkaProperties),
		dependency.NewIdGenerator(),
	)

	ratingInteractor.GetIncomingMessage()
}


