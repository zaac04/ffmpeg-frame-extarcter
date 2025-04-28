package worker

import (
	"encoding/json"
	"fmt"
)

func doSaveResults(w *Worker) error {

	id := w.Job.Details.InterviewId
	results := w.Job.Details.Results
	data, _ := json.Marshal(results)
	fmt.Println(id, string(data))
	return nil
}
