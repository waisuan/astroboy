package datapipeline

import (
	k "astroboy/internal/kafka"
	s "astroboy/internal/sqs"
	"github.com/segmentio/kafka-go"
	"sync"
)

var Memory sync.Map

func Run() {
	sqsCli := s.NewSqsCli()

	sqsMailbox := make(chan string)
	defer close(sqsMailbox)

	go sqsCli.ReceiveMessage(sqsMailbox)

	kafkaCli := k.NewKafkaCli()

	err := kafkaCli.CreateTopic()
	if err != nil {
		panic(err)
	}

	kafkaMailbox := make(chan string)
	defer close(kafkaMailbox)

	go kafkaCli.ConsumeMessage(kafkaMailbox)

	err = kafkaCli.ProduceMessage(kafka.Message{
		Key:   []byte("Some-Key"),
		Value: []byte("Hello, Universe!"),
	})
	if err != nil {
		panic(err)
	}

	for {
		kafkaMessage := <-kafkaMailbox

		err = sqsCli.SendMessage(kafkaMessage)
		if err != nil {
			panic(err)
		}

		sqsMessage := <-sqsMailbox
		Memory.Store("LatestMessage", sqsMessage)
	}
}
