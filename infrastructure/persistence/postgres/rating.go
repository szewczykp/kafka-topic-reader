package postgres

import (
	"github.com/go-pg/pg/v9"
	"kafka_topic_reader/domain/entity"
	"kafka_topic_reader/domain/repository"
)

type ratingRepository struct {
	db *pg.DB
}

func (r *ratingRepository) Save(rating *entity.Rating) error {
	err := r.db.Insert(rating)

	if err != nil {
		return err
	}

	return nil
}

func NewRatingRepository(db *pg.DB) repository.RatingRepository {
	return &ratingRepository{db: db}
}
