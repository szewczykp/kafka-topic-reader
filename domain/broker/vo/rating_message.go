package vo

import (
	"encoding/json"
	"fmt"
)

type RatingMessage struct {
	Id       string `json:"id"`
	RecipeId string `json:"recipe_id"`
	Value    int8   `json:"value"`
}

func CreateRatingMessageFromJSON(data []byte) (RatingMessage, error) {
	var ratingMessage RatingMessage
	err := json.Unmarshal(data, &ratingMessage)

	if err != nil {
		return RatingMessage{}, err
	}

	fmt.Println("Raw message was parsed successfully")
	return newRatingCreator(ratingMessage)
}

func newRatingCreator(rc RatingMessage) (RatingMessage, error) {
	if (rc == RatingMessage{}) {
		return RatingMessage{}, fmt.Errorf("rating message is empty")
	}
	if rc.Id == "" || rc.Value == int8(0) || rc.RecipeId == "" {
		return RatingMessage{}, fmt.Errorf("rating message has an invalid format")
	}

	return rc, nil
}
