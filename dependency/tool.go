package dependency

import (
	"kafka_topic_reader/domain/tool"
	tool2 "kafka_topic_reader/infrastructure/tool"
)

func NewIdGenerator() tool.IdGenerator {
	return tool2.NewUuidGenerator()
}