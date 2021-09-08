package interactor

import (
	"kafka_topic_reader/domain/broker"
	"kafka_topic_reader/domain/entity"
	"kafka_topic_reader/domain/repository"
	"kafka_topic_reader/domain/tool"
	"log"
)

type RatingInteractor interface {
	create(rating *entity.Rating)
	GetIncomingMessage() error
}

type ratingInteractor struct {
	ratingRepository repository.RatingRepository
	ratingConsumer   broker.RatingConsumer
	idGenerator      tool.IdGenerator
}

// GetIncomingMessage TODO: transfer incoming message to new struct. Save message to DB
func (ri *ratingInteractor) GetIncomingMessage() error {
	ratingMessage, err := ri.ratingConsumer.GetRatingMessage()
	rating := entity.Rating(ratingMessage)

	if err == nil {
		ri.create(&rating)
	} else {
		return err
	}

	return nil
}

func (ri *ratingInteractor) create(rating *entity.Rating) {
	err := ri.ratingRepository.Save(rating)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}

func NewRatingInteractor(
	ratingRepository repository.RatingRepository,
	ratingConsumer broker.RatingConsumer,
	idGenerator tool.IdGenerator,
) RatingInteractor {
	return &ratingInteractor{
		ratingRepository: ratingRepository,
		ratingConsumer:   ratingConsumer,
		idGenerator:      idGenerator,
	}
}
