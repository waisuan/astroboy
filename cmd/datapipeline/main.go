package main

import (
	"astroboy/internal/cache"
	k "astroboy/internal/kafka"
	s "astroboy/internal/sqs"
	"fmt"
	"time"
)

func main() {
	cacheCli := cache.NewCache()

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

	for {
		fmt.Println("Waiting for messages...")

		kafkaMessage := <-kafkaMailbox

		err = sqsCli.SendMessage(kafkaMessage)
		if err != nil {
			panic(err)
		}

		sqsMessage := <-sqsMailbox

		err := cacheCli.Set("latest_message", sqsMessage, 1*time.Hour)
		if err != nil {
			panic(err)
		}
	}
}
