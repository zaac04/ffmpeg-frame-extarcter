package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)

type Kafka struct {
	client sarama.Client
}

type KafkaTopic struct {
	Name              []string
	ReplicationFactor int16
	NumPartitions     int32
}

func CreateClient(kafkaEndpoint []string) (*Kafka, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	fmt.Println("endpoint here", kafkaEndpoint)
	client, err := sarama.NewClient(kafkaEndpoint, config)
	fmt.Println("rrr here", err, kafkaEndpoint)
	return &Kafka{client: client}, err
}

func Ping(kafkaEndpoint []string) error {
	_, err := CreateClient(kafkaEndpoint)
	return err
}

func CreateTopic(kafkaEndpoint []string, TopicDetails *KafkaTopic) (err error) {
	fmt.Println(kafkaEndpoint)
	kafka, err := CreateClient(kafkaEndpoint)

	if err != nil {
		fmt.Println("dadasda", err)
		return
	}

	admin, err := sarama.NewClusterAdminFromClient(kafka.client)

	if err != nil {
		fmt.Println("asdsasada", err)
		return
	}

	current_topics, err := admin.ListTopics()

	for _, topic := range TopicDetails.Name {
		if _, ok := current_topics[topic]; ok {
			continue
		}
		err = admin.CreateTopic(
			topic,
			&sarama.TopicDetail{
				NumPartitions:     TopicDetails.NumPartitions,
				ReplicationFactor: TopicDetails.ReplicationFactor,
			},
			false)
	}

	return err
}
