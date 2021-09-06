package repository

import "kafka_topic_reader/domain/entity"

type RatingRepository interface {
	Save(*entity.Rating) error
}