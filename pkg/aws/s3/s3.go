package s3

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DownloadFromS3(bucket_name string, key string) (location io.ReadCloser, err error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}
	fmt.Println("Downloading file:", bucket_name, key)
	resp, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucket_name,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("Download Completed:", bucket_name, key)
	return resp.Body, err

}

func getClient() (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg)
	return client, err
}

func WriteToS3(bucket_name string, key string, body string) (location string, err error) {
	client, err := getClient()
	if err != nil {
		return "", err
	}
	fmt.Println("Uploading file:", bucket_name, key)
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket_name,
		Key:    &key,
		Body:   bytes.NewReader([]byte(body)),
	})
	if err != nil {
		return "", err
	}

	fmt.Println("Upload Completed:", bucket_name, key)

	return fmt.Sprintf("s3://%s/%s", bucket_name, key), err
}
