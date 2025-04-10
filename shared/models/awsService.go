package models

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type S3Service struct {
	Client *s3.Client
	Bucket string
}

type SQSService struct {
	Client *sqs.Client
	Queue  string
}
