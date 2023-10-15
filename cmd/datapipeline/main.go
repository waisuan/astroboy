package main

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/model"
	"encoding/json"
	"log"
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
		log.Println("Waiting for messages...")

		kafkaMessage := <-kafkaMailbox

		err = deps.SqsCli.SendMessage(kafkaMessage)
		if err != nil {
			panic(err)
		}

		sqsMessage := <-sqsMailbox

		var m model.Message
		err = json.Unmarshal([]byte(sqsMessage), &m)
		if err != nil {
			panic(err)
		}

		switch m.T {
		case "user":
			err := deps.CacheCli.Set(m.Key, m.Payload, 1*time.Hour)
			if err != nil {
				panic(err)
			}
		default:
			err := deps.CacheCli.Set("latest_message", m.Payload, 1*time.Hour)
			if err != nil {
				panic(err)
			}
		}
	}
}
