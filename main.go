package main

import (
	"context"
	"fmt"
	"nishchal/configs"
	"nishchal/internal/initializers"
	"nishchal/internal/queue"
	"nishchal/internal/utilities"
	"nishchal/internal/worker"
	"sync"
)

func init() {
	fmt.Println("Nischal Starting Up...")
	initializers.Init()
}

func main() {

	var wg sync.WaitGroup
	workload_queue := make(chan queue.Message, 5)
	results_queue := make(chan queue.Message, 5)

	queue, err := queue.CreateNewQueue([]string{configs.APP.KafkaEndpoint})
	utilities.FailOnError(err)

	go queue.Receive(context.Background(), workload_queue, []string{"start-analysis"}, "analysis")

	//spin up worker pods for
	for i := range worker.MaxWorkers {
		wg.Add(1)
		go worker.CreateWorker(workload_queue, &wg).StartWorking(i)
	}

	go queue.Receive(context.Background(), results_queue, []string{"results-topic"}, "results")

	for i := range worker.MaxWorkers {
		wg.Add(1)
		go worker.CreateWorker(results_queue, &wg).StartWorking(i)
	}

	fmt.Println("Nischal Up and Running...")

	wg.Wait()
}
