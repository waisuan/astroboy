package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const QUEUE_URL = "http://localhost:4566/000000000000/test-q"

type SqsCli struct {
	conn *sqs.SQS
}

func NewSqsCli() *SqsCli {
	cfg := aws.Config{
		Region:   aws.String("eu-west-1"),
		Endpoint: aws.String("http://localhost:4566/"),
	}

	sess := session.Must(session.NewSession(&cfg))

	return &SqsCli{
		conn: sqs.New(sess),
	}
}

func (s *SqsCli) SendMessage(msg string) error {
	message := &sqs.SendMessageInput{
		QueueUrl:    aws.String(QUEUE_URL),
		MessageBody: aws.String(msg),
	}

	_, err := s.conn.SendMessage(message)
	return err
}

func (s *SqsCli) ReceiveMessage(mailbox chan string) error {
	param := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(QUEUE_URL),
	}

	for {
		m, err := s.conn.ReceiveMessage(param)
		if err != nil {
			return err
		}

		for _, v := range m.Messages {
			mailbox <- *v.Body

			_, err = s.conn.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(QUEUE_URL),
				ReceiptHandle: v.ReceiptHandle,
			})
		}
	}
}
