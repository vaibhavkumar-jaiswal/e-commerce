// AWS S3 and SQS operations
package services

import (
	"context"
	"e-commerce/shared/models"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Service models.S3Service

func NewS3Service(bucket string) *models.S3Service {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Unable to load AWS config: %v", err)
	}
	return &models.S3Service{
		Client: s3.NewFromConfig(cfg),
		Bucket: bucket,
	}
}

func (service *S3Service) UploadFile(filePath string, key string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = service.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &service.Bucket,
		Key:    &key,
		Body:   file,
		ACL:    types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return err
	}
	log.Printf("File uploaded to S3: %s", key)
	return nil
}

func (service *S3Service) ListFiles() ([]string, error) {
	output, err := service.Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &service.Bucket,
	})
	if err != nil {
		return nil, err
	}

	var files []string
	for _, obj := range output.Contents {
		files = append(files, *obj.Key)
	}
	return files, nil
}
