package utilities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"path"
	"strings"
)

func FailOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckErrorAndExit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func UnMarshalJson(data []byte, dst interface{}) error {
	r := bytes.NewReader(data)
	decoder := json.NewDecoder(r)
	return decoder.Decode(&dst)
}

func ParseS3Url(s3url string) (bucket string, key string, base string, err error) {
	parsed, err := url.Parse(s3url)
	if err != nil {
		return
	}
	base = path.Base(parsed.Path)

	switch parsed.Scheme {
	case "s3":
		bucket = parsed.Host
		key = strings.TrimLeft(parsed.Path, "/")
	case "https":
		hostParts := strings.Split(parsed.Host, ".")
		if len(hostParts) < 4 || hostParts[1] != "s3" {
			return "", "", "", fmt.Errorf("not a valid S3 HTTPS URL")
		}
		bucket = hostParts[0]
		key = strings.TrimLeft(parsed.Path, "/")
	default:
		return "", "", "", fmt.Errorf("unsupported URL scheme: %s", parsed.Scheme)
	}

	return bucket, key, base, nil

}
