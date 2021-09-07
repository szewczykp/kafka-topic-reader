package postgres

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"kafka_topic_reader/config"
	"kafka_topic_reader/domain/entity"
	"kafka_topic_reader/domain/tool"
	tool3 "kafka_topic_reader/infrastructure/tool"
	"testing"
)

type RatingRepositoryTestSuite struct {
	suite.Suite
	db          *pg.DB
	idGenerator tool.IdGenerator
}

func (suite *RatingRepositoryTestSuite) SetupTest() {
	host, port, user, password, database := config.GetDatabaseConfig()

	suite.db = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		User:     user,
		Password: password,
		Database: database,
	})
	suite.idGenerator = tool3.NewUuidGenerator()
}

func (suite *RatingRepositoryTestSuite) TestSave() {
	rating := &entity.Rating{
		Id:       suite.idGenerator.Generate(),
		RecipeId: suite.idGenerator.Generate(),
		Value:    5,
	}

	underTest := ratingRepository{db: suite.db}

	err := underTest.Save(rating)

	assert.Nil(suite.T(), err)
}

func TestRatingRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RatingRepositoryTestSuite))
}
