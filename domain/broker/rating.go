package broker

type RatingMessage struct {

}

type RatingConsumer interface {
	GetRatingMessage()
}
