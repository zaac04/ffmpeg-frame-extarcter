package configs

import (
	"fmt"
	"log"
	"nishchal/pkg/env"
)

var APP appConfig

type appConfig struct {
	QueueEndpoint   string   `env:"QUEUE_ENDPOINT"`
	FileDestination string   `env:"FILE_DESTINATION"`
	TopicName       []string `env:"TOPIC_NAME"`
	KafkaEndpoint   string   `env:"KAFKA_ENDPOINT"`
}

func LoadEnv(env_file string) {
	err := env.Load_env(env_file, &APP)
	if err != nil {
		log.Fatalf("failed to load env: %v", err)
	}

	fmt.Printf("ENV's Loaded and validation was success!\n")

	fmt.Println(APP)
}
