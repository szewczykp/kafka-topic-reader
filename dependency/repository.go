package dependency

import (

	"github.com/go-pg/pg/v9"
	"kafka_topic_reader/domain/repository"
	"kafka_topic_reader/infrastructure/persistence/postgres"
)

func NewRatingRepository(db interface{}) repository.RatingRepository {
	switch connection := db.(type) {
	case *pg.DB:
		return postgres.NewRatingRepository(connection)
	}

	return nil
}