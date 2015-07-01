package ls

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// FileInfo is the information of File
type FileInfo struct {
	Name         string
	LastModified time.Time
	Owner        string
	Size         uint64
}

// Config is the configuration for list files
type Config struct {
	Region string
	Bucket string
}

// List files
func List(file string, config *Config) ([]*FileInfo, error) {
	client := s3.New(&aws.Config{Region: config.Region})

	if len(file) > 0 && file[0] == '/' {
		file = file[1:]
	}

	// Obtem os arquivos diretorio
	params := &s3.ListObjectsInput{
		Bucket: aws.String(config.Bucket),
		Prefix: aws.String(file),
	}
	response, err := client.ListObjects(params)

	if err != nil {
		return nil, err
	}

	if *response.IsTruncated {
		return nil, fmt.Errorf(
			"O limite máximo de arquivos no backup por cliente que é de %d foi ultrapassado.",
			*response.MaxKeys)
	}

	files := make([]*FileInfo, len(response.Contents))
	for i, obj := range response.Contents {
		files[i] = &FileInfo{
			Name:         *obj.Key,
			LastModified: *obj.LastModified,
			Owner:        *obj.Owner.DisplayName,
			Size:         uint64(*obj.Size),
		}
	}

	return files, nil
}
