package kafka

import (
	"github.com/IBM/sarama"
)

func (k *Kafka) ConsumeFromTopic(group_id string) (sarama.ConsumerGroup, error) {
	return sarama.NewConsumerGroupFromClient(group_id, k.client)
}
