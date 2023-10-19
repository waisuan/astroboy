package workers

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/model"
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gocraft/work"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func (c *Context) PublishFakeData(_ *work.Job) error {
	log.Println("Running publish_fake_data job...")

	deps := dependencies.Init()

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

	err := deps.KafkaCli.ProduceMessage(kafka.Message{
		Value: mj,
	})

	return err
}
