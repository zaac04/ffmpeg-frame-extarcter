package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
)

type SendMsgBody struct {
	Topic string
	Value interface{}
}

func (k *Kafka) PublishMessage(msgBody *SendMsgBody) error {
	producer, err := sarama.NewSyncProducerFromClient(k.client)
	defer producer.Close()
	if err != nil {
		return err
	}
	fmt.Println("Topic", msgBody.Topic, "Value", msgBody.Value)

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: msgBody.Topic,
		Value: JsonEncoder{data: msgBody.Value},
	})

	return err
}

type JsonEncoder struct {
	data any
	json []byte
}

func (j JsonEncoder) Encode() (data []byte, err error) {
	j.json, err = json.Marshal(j.data)
	fmt.Println("json", string(j.json))
	return j.json, err
}

func (j JsonEncoder) Length() int {
	return len(j.json)
}
