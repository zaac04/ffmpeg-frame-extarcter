package worker

import (
	"nishchal/internal/queue"
	"sync"
)

var MaxWorkers = 3
var Workers AllWorker

type AllWorker struct {
	Workers []Worker
}

type Worker struct {
	Id     string
	Status string
	Queue  chan queue.Message
	Wg     *sync.WaitGroup
	Job    *Job
}

type Job struct {
	Type     string    `json:"type"`
	Details  JobDetail `json:"detail"`
	HandleId string    `json:"queue_handle_id"`
}

type JobDetail struct {
	InterviewId    string      `json:"interview_id"`
	SourceLocation string      `json:"src_location"`
	Results        interface{} `json:"results,omitempty"`
}
