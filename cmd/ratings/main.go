package main

import (
	"fmt"
	"kafka_topic_reader/config"
	"kafka_topic_reader/dependency"
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

	//ratingInteractor := interactor.NewRatingInteractor(
	//	dependency.NewRatingRepository(db),
	//	dependency.NewIdGenerator(),
	//)

	ratingConsumer := dependency.NewRatingConsumer(kafkaProperties)
	ratingConsumer.GetRatingMessage()

}


