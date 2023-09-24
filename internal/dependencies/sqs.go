package dependencies

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SqsCli struct {
	cfg  *Config
	conn *sqs.SQS
}

func NewSqsCli(cfg *Config) *SqsCli {
	awsCfg := aws.Config{
		Region:   aws.String(cfg.AwsRegion),
		Endpoint: aws.String(cfg.AwsSqsEndpoint),
	}

	sess := session.Must(session.NewSession(&awsCfg))

	return &SqsCli{
		cfg:  cfg,
		conn: sqs.New(sess),
	}
}

func (s *SqsCli) SendMessage(msg string) error {
	message := &sqs.SendMessageInput{
		QueueUrl:    aws.String(s.cfg.SqsQueueUrl),
		MessageBody: aws.String(msg),
	}

	_, err := s.conn.SendMessage(message)
	return err
}

func (s *SqsCli) ReceiveMessage(mailbox chan string) error {
	param := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(s.cfg.SqsQueueUrl),
	}

	for {
		m, err := s.conn.ReceiveMessage(param)
		if err != nil {
			return err
		}

		for _, v := range m.Messages {
			mailbox <- *v.Body

			_, err = s.conn.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(s.cfg.SqsQueueUrl),
				ReceiptHandle: v.ReceiptHandle,
			})
		}
	}
}
