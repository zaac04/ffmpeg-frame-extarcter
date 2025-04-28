package initializers

import (
	"fmt"
	"nishchal/configs"
	"nishchal/internal/utilities"
	"nishchal/pkg/kafka"
)

func Init() {

	configs.LoadEnv(".env")

	endpoint := []string{configs.APP.KafkaEndpoint}

	utilities.FailOnError(kafka.Ping(endpoint))
	fmt.Println("Kafka Running at:", endpoint)
	fmt.Println(configs.APP.TopicName)
	utilities.FailOnError(
		kafka.CreateTopic(endpoint, &kafka.KafkaTopic{
			Name:              configs.APP.TopicName,
			NumPartitions:     10,
			ReplicationFactor: 1,
		}))
}
