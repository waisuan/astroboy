package main

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/model"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	deps := dependencies.Init()

	err := deps.KafkaCli.CreateTopic()
	if err != nil {
		log.Fatal(err)
	}

	kafkaMailbox := make(chan string)
	defer close(kafkaMailbox)

	go func() {
		err := deps.KafkaCli.ConsumeMessage(kafkaMailbox)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = deps.KafkaCli.ProduceMessage(kafka.Message{
		Value: fakeUser(),
	})
	if err != nil {
		log.Fatal(err)
	}

	kafkaMessage := <-kafkaMailbox
	fmt.Println(kafkaMessage)

	err = deps.KafkaCli.ProduceMessage(kafka.Message{
		Value: fakeUser(),
	})
	if err != nil {
		log.Fatal(err)
	}

	kafkaMessage = <-kafkaMailbox
	fmt.Println(kafkaMessage)
}

func fakeUser() []byte {
	username := gofakeit.Name()

	startDate, _ := time.Parse("2006-01-02", "1980-01-01")
	endDate, _ := time.Parse("2006-01-02", "2010-01-01")
	dob := gofakeit.DateRange(startDate, endDate)

	user := model.User{
		Username:    username,
		Email:       gofakeit.Email(),
		DateOfBirth: dob.Format("2006-01-02"),
	}
	uj, _ := json.Marshal(user)

	m := model.Message{
		T:       "user",
		Key:     username,
		Payload: string(uj),
	}
	mj, _ := json.Marshal(m)

	return mj
}
