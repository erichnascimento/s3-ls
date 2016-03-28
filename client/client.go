package client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Config is the configuration for list files
type Config struct {
	Region string
	Bucket string
}

// Client is the s3 client
type Client struct {
	*s3.Client
	Config *Config
}

// NewClient return a new s3 Client
func NewClient(config *Config) (*Client, error) {
	client := &Client{
		s3.New(&aws.Config{Region: config.Region}),
		config,
	}

	return client, nil
}
