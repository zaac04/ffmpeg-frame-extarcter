package filemanager

import (
	"fmt"
	"nishchal/configs"
	"nishchal/internal/enums"
	"nishchal/internal/utilities"
	"nishchal/pkg/aws/s3"
	"os"
	"path"
)

func DownloadToLocalFromS3(Source string, InterviewId string) (location string, err error) {

	destination := configs.APP.FileDestination
	bucket, key, filename, err := utilities.ParseS3Url(Source)
	if err != nil {
		return "", err
	}

	resp, err := s3.DownloadFromS3(bucket, key)
	if err != nil {
		return "", err
	}

	defer resp.Close()

	folder := path.Join(destination, enums.SRC, InterviewId)
	location = path.Join(folder, filename)
	err = os.MkdirAll(folder, 0766)

	if err != nil {
		return "", err
	}
	fmt.Println("Created Folder:", folder)

	outfile, err := os.Create(location)
	if err != nil {
		return "", err
	}
	defer outfile.Close()

	fmt.Println("Writing to file:", location)
	_, err = outfile.ReadFrom(resp)
	if err != nil {
		return "", err
	}
	fmt.Println("Writing to file:", location, "Completed")
	return
}

func WriteFiletoS3(bucket_name string, content string, InterviewId string) (location string, err error) {

	bucket, _, _, err := utilities.ParseS3Url(bucket_name)
	if err != nil {
		return "", err
	}

	fmt.Println("Writing to S3 in progress", bucket)

	return s3.WriteToS3(bucket, fmt.Sprintf("analyzed-data/%s.json", InterviewId), content)
}
