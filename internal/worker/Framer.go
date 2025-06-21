package worker

import (
	"context"
	"fmt"
	"log"
	"nishchal/internal/filemanager"
	"strings"
	"time"

	"google.golang.org/genai"
)

func doMakeFrameJob(w *Worker) error {

	src_location, err := filemanager.DownloadToLocalFromS3(
		w.Job.Details.SourceLocation,
		w.Job.Details.InterviewId,
	)

	if err != nil {
		return err
	}

	dest := path.Join(
		configs.APP.FileDestination,
		enums.Frames,
		w.Job.Details.InterviewId,
	)
	ffmpeg.MakeFrames(src_location, dest)

	q, err := queue.CreateNewQueue([]string{configs.APP.KafkaEndpoint})

	if err != nil {
		return err
	}

	err = q.Send(queue.SendMsgBody{
		Topic: "ml-topic",
		Value: Job{
			Type: "Analyze",
			Details: JobDetail{
				InterviewId:    w.Job.Details.InterviewId,
				SourceLocation: dest,
			},
		},
	})

	filemanager.WriteFiletoS3(w.Job.Details.SourceLocation, cleaned, w.Job.Details.InterviewId)

	return err
}
