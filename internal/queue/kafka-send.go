package queue

import "nishchal/pkg/kafka"

type Kafka struct {
	client kafka.Kafka
}

func (k *Kafka) Send(m SendMsgBody) error {
	return k.client.PublishMessage(&kafka.SendMsgBody{
		Topic: m.Topic,
		Value: m.Value,
	})
}
