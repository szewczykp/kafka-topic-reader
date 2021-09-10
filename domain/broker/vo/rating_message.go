package vo

import (
	"encoding/json"
	"fmt"
	"kafka_topic_reader/domain/broker"
)

func CreateRatingMessageFromJSON(data []byte) (broker.RatingMessage, error) {
	var ratingMessage broker.RatingMessage
	err := json.Unmarshal(data, &ratingMessage)

	if err != nil {
		return broker.RatingMessage{}, err
	}

	fmt.Println("Raw message was parsed successfully")
	return newRatingCreator(ratingMessage)
}

func newRatingCreator(rc broker.RatingMessage) (broker.RatingMessage, error) {
	if (rc == broker.RatingMessage{}) {
		return broker.RatingMessage{}, fmt.Errorf("rating message is empty")
	}
	if rc.Id == "" || rc.Value == int8(0) || rc.RecipeId == "" {
		return broker.RatingMessage{}, fmt.Errorf("rating message has an invalid format")
	}

	return rc, nil
}
