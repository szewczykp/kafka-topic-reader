package interactor

import (
	"kafka_topic_reader/domain/broker"
	"kafka_topic_reader/domain/entity"
	"kafka_topic_reader/domain/repository"
	"kafka_topic_reader/domain/tool"
)

type RatingInteractor interface {
	create(rating *entity.Rating) error
	GetIncomingMessage()
}

type ratingInteractor struct {
	ratingRepository repository.RatingRepository
	ratingConsumer broker.RatingConsumer
	idGenerator tool.IdGenerator
}

func (ri *ratingInteractor) create(rating *entity.Rating) error {
	rating.Id = ri.idGenerator.Generate()
	err := ri.ratingRepository.Save(rating)

	if err != nil {
		return err
	}
	return nil
}

// GetIncomingMessage TODO: transfer incoming message to new struct. Save message to DB
func (ri *ratingInteractor) GetIncomingMessage() {
	ri.ratingConsumer.GetRatingMessage()
}

func NewRatingInteractor(
	ratingRepository repository.RatingRepository,
	ratingConsumer broker.RatingConsumer,
	idGenerator tool.IdGenerator,
	) RatingInteractor{
	return &ratingInteractor{
		ratingRepository: ratingRepository,
		ratingConsumer: ratingConsumer,
		idGenerator: idGenerator,
	}
}