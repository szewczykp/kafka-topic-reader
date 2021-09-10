package broker

import "kafka_topic_reader/domain/broker/vo"

type RatingConsumer interface {
	GetRatingMessage() (vo.RatingMessage, error)
}
