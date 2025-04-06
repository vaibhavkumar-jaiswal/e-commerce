package services

import (
	"context"
	"e-commerce/shared/models"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSService models.SQSService

func NewSQSService(queueURL string) *SQSService {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Printf("Unable to load AWS config: %v", err)
	}
	return &SQSService{
		Client: sqs.NewFromConfig(cfg),
		Queue:  queueURL,
	}
}

func (s *SQSService) SendMessage(message string) error {
	_, err := s.Client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    &s.Queue,
		MessageBody: &message,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Message sent: %s", message)
	return nil
}

func (s *SQSService) ReceiveMessages() ([]string, error) {
	output, err := s.Client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:            &s.Queue,
		MaxNumberOfMessages: 10,
		WaitTimeSeconds:     10,
	})
	if err != nil {
		return nil, err
	}

	var messages []string
	for _, msg := range output.Messages {
		messages = append(messages, *msg.Body)
	}
	return messages, nil
}
