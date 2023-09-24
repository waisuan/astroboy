package main

import (
	"astroboy/internal/dependencies"
	"fmt"
	"time"
)

func main() {
	deps := dependencies.Init()

	sqsMailbox := make(chan string)
	defer close(sqsMailbox)

	go deps.SqsCli.ReceiveMessage(sqsMailbox)

	err := deps.KafkaCli.CreateTopic()
	if err != nil {
		panic(err)
	}

	kafkaMailbox := make(chan string)
	defer close(kafkaMailbox)

	go deps.KafkaCli.ConsumeMessage(kafkaMailbox)

	for {
		fmt.Println("Waiting for messages...")

		kafkaMessage := <-kafkaMailbox

		err = deps.SqsCli.SendMessage(kafkaMessage)
		if err != nil {
			panic(err)
		}

		sqsMessage := <-sqsMailbox

		err := deps.CacheCli.Set("latest_message", sqsMessage, 1*time.Hour)
		if err != nil {
			panic(err)
		}
	}
}
