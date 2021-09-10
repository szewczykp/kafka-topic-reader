package kafka_broker

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"kafka_topic_reader/config"
	"kafka_topic_reader/domain/broker/vo"
	"reflect"
	"strings"
	"testing"
	"time"
)

type RatingConsumerTestSuite struct {
	suite.Suite
	kafkaTestCluster *testcontainers.LocalDockerCompose
	consumer         *kafka.Consumer
}

func (suite *RatingConsumerTestSuite) SetupSuite() {
	host, port, consumerTopicName, groupId := config.GetConsumerConfig()

	suite.kafkaTestCluster = testcontainers.NewLocalDockerCompose(
		[]string{"../../test-docker-compose.yml"},
		strings.ToLower(uuid.New().String()),
	)

	suite.kafkaTestCluster.WithCommand([]string{"up", "-d", "zookeeper", "kafka"}).Invoke()
	time.Sleep(5 * time.Second)

	var err error
	suite.consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", host, port),
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	subscriberError := suite.consumer.Subscribe(consumerTopicName, nil)

	if subscriberError != nil {
		panic(subscriberError)
	}
}

func (suite *RatingConsumerTestSuite) TearDownSuite() {
	suite.consumer.Close()
	suite.kafkaTestCluster.Down()
}

func (suite *RatingConsumerTestSuite) TestReceivingRatingMessageFromKafka() {
	//given
	ratingMessage := vo.RatingMessage{
		Id:       "7faaa7f2-1079-11ec-82a8-0242ac130003",
		RecipeId: "c0d7c99e-1079-11ec-82a8-0242ac130003",
		Value:    5,
	}

	kafkaConsumerUnderTest := ratingConsumer{consumer: suite.consumer}
	producedMessage, _ := json.Marshal(ratingMessage)

	//when
	produceKafkaMessage(string(producedMessage))
	consumedMessage, _ := kafkaConsumerUnderTest.GetRatingMessage()

	//then
	if !reflect.DeepEqual(ratingMessage, consumedMessage) {
		suite.T().Fatalf("Consumed messages are not equal to produced messages. %v != %v", consumedMessage, producedMessage)
	}
}

func produceKafkaMessage(message string) {
	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", "localhost", "9092"),
	})
	if err != nil {
		panic(err)
	}
	defer kafkaProducer.Close()

	fmt.Printf("Producing messages into kafka...\n")

	topic := "test-topic"

	deliveryChan := make(chan kafka.Event)

	kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message [%s] to topic %s [%d] at offset %v\n",
			message, *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
}

func TestRatingRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RatingConsumerTestSuite))
}
