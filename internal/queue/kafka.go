package queue

import (
	"nishchal/pkg/kafka"
)

func NewKafkaClient(endpoint []string) (*Kafka, error) {
	client, err := kafka.CreateClient(endpoint)
	return &Kafka{
		client: *client,
	}, err
}
