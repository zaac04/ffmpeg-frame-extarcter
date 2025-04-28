package ffmpeg

import (
	"bytes"
	"fmt"
	"math"
	"nishchal/internal/utilities"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

func MakeFrames(src_file string, dest string) {

	var wg sync.WaitGroup
	now := time.Now()

	//timing
	concurrency := 1
	duration, err := getDuration(src_file)

	if err != nil {
		fmt.Println(err)
		return
	}

	toalDurationInSeconds := int(math.Ceil(duration / float64((concurrency))))
	NumofMinutes := toalDurationInSeconds / 60
	NumofSeconds := toalDurationInSeconds % 60
	fmt.Println("Total duration estimated is: ", NumofMinutes, "mins", NumofSeconds, "secs")

	err = os.MkdirAll(dest, 0766)
	fmt.Println("Created Output Directory", dest)
	if err != nil {
		return
	}

	fmt.Println("Started Converting Video to frames", src_file)
	for job := range NumofMinutes {
		wg.Add(1)
		go func(frame int) {
			defer wg.Done()
			executeFFMPEG(frame, 60, src_file, dest)
		}(job)
	}

	wg.Add(1)
	go func(frame int) {
		defer wg.Done()
		executeFFMPEG(frame, NumofSeconds, src_file, dest)
	}(NumofMinutes)

	fmt.Println("Converted videos to frame", dest)
	wg.Wait()
	fmt.Print(time.Since(now))
}

func getDuration(file_name string) (duration float64, err error) {

	stdout, _ := executeCmd(
		"ffprobe",
		"-v", "quiet",
		"-print_format", "json",
		"-show_format", file_name,
	)
	var Output FFProbeOutput
	utilities.UnMarshalJson([]byte(stdout), &Output)
	duration, err = strconv.ParseFloat(Output.Format.Duration, 64)
	if err != nil {
		return 0.0, err
	}
	return duration, nil
}

func executeCmd(command string, arg ...string) (stdout string, stderr string) {

	cmd := exec.Command(command, arg...)
	var stdoutBuffer, stderrBuffer bytes.Buffer
	cmd.Stdout = &stdoutBuffer
	cmd.Stderr = &stderrBuffer

	if err := cmd.Run(); err != nil {
		fmt.Println("error", err)
		fullLog := stdoutBuffer.String() + stderrBuffer.String()
		fmt.Println(fullLog)
	}

	return stdoutBuffer.String(), stderrBuffer.String()
}

func executeFFMPEG(StartTime int, duration int, file string, destination string) {
	startTime := StartTime * 60
	outputPattern := fmt.Sprintf("%s/%%03d.jpg", destination)
	executeCmd(
		"ffmpeg",
		// "-hwaccel", "cuda",
		"-ss", fmt.Sprintf("%d", startTime),
		"-t", strconv.Itoa(duration),
		"-i", file,
		"-vf", "fps=1",
		"-q:v", "10",
		"-start_number", strconv.Itoa(startTime),
		outputPattern,
	)
}
