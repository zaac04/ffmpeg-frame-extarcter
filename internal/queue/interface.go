package queue

import (
	"context"
)

type Message struct {
	Body []byte
}

type SendMsgBody struct {
	Topic string
	Value interface{}
}

type Handle string

type Queue interface {
	Send(msg SendMsgBody) error
	Receive(ctx context.Context, workload_queue chan<- Message, topic []string, grp_id string) error
}

func CreateNewQueue(endpoint []string) (Queue, error) {
	return NewKafkaClient(endpoint)
}
