package interactor

import (
	"kafka_topic_reader/domain/entity"
	"kafka_topic_reader/domain/repository"
	"kafka_topic_reader/domain/tool"
)

type RatingInteractor interface {
	Create(rating *entity.Rating) error
}

type ratingInteractor struct {
	ratingRepository repository.RatingRepository
	idGenerator tool.IdGenerator
}

func (ri *ratingInteractor) Create(rating *entity.Rating) error {
	rating.Id = ri.idGenerator.Generate()
	err := ri.ratingRepository.Save(rating)

	if err != nil {
		return err
	}
	return nil
}

func NewRatingInteractor(
	ratingRepository repository.RatingRepository,
	idGenerator tool.IdGenerator,
	) RatingInteractor{
	return &ratingInteractor{
		ratingRepository: ratingRepository,
		idGenerator: idGenerator,
	}
}