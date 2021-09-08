package broker

type RatingMessage struct {
	Id       string `json:"id"`
	RecipeId string `json:"recipe_id"`
	Value    int8   `json:"value"`
}

type RatingConsumer interface {
	GetRatingMessage() (RatingMessage, error)
}
