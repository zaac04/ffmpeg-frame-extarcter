package queue

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
)

func (k *Kafka) Receive(ctx context.Context, workload_queue chan<- Message, topics []string, grp_id string) error {
	consumer, err := k.client.ConsumeFromTopic(grp_id)
	if err != nil {
		return err
	}

	fmt.Println("Kafka listening on topic", topics)
	err = consumer.Consume(context.Background(), topics, &ConsumerGrpHandler{
		MsgCh: workload_queue,
	})

	if err != nil {
		return err
	}

	return nil

}

func (h *ConsumerGrpHandler) ConsumeClaim(
	session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		h.MsgCh <- Message{
			Body: message.Value,
		}
		session.MarkMessage(message, "")
	}
	return nil
}

type ConsumerGrpHandler struct {
	MsgCh chan<- Message
}

func (h *ConsumerGrpHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (h *ConsumerGrpHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}
