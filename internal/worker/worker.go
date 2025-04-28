package worker

import (
	"fmt"
	"nishchal/internal/enums"
	"nishchal/internal/queue"
	"nishchal/internal/utilities"
	"sync"

	"github.com/google/uuid"
)

func CreateWorker(ch chan queue.Message, wg *sync.WaitGroup) *Worker {
	worker := Worker{
		Id:     uuid.NewString(),
		Status: enums.WorkerIdle,
		Queue:  ch,
		Wg:     wg,
	}
	Workers.Workers = append(Workers.Workers, worker)
	return &worker
}

func (w *Worker) StartWorking(i int) {

	fmt.Println("Worker", i, w.Id, "Started Up")

	for msg := range w.Queue {
		w.decodeJob(msg)
		w.findHandlerAndStart()
	}

	defer w.Wg.Done()
}

func (w *Worker) decodeJob(msg queue.Message) {
	job := &Job{}
	utilities.UnMarshalJson(msg.Body, job)
	w.Job = job
	fmt.Println("hewjh")
	fmt.Printf("%+v\n", w.Job)
}

func (w *Worker) findHandlerAndStart() (err error) {

	switch w.Job.Type {
	case "Framer":
		return w.doJob(doMakeFrameJob)

	case "Analyzer-Results":
		return w.doJob(doSaveResults)

	default:
		return w.doJob(func(w *Worker) error {
			fmt.Println("default")
			return err
		})

	}
}

func (w *Worker) doJob(fn func(w *Worker) error) error {
	return fn(w)
}
