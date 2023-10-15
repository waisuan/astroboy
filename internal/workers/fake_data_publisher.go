package workers

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/model"
	"encoding/json"
	"github.com/gocraft/work"
	"github.com/segmentio/kafka-go"
	"log"
)

func (c *Context) PublishFakeData(job *work.Job) error {
	log.Println("Running publish_fake_data job...")

	deps := dependencies.Init()

	user := model.User{
		Username:    "esia",
		Email:       "e-sia@outlook.com",
		DateOfBirth: "12/11/1991",
	}
	uj, _ := json.Marshal(user)

	m := model.Message{
		T:       "user",
		Key:     "esia",
		Payload: string(uj),
	}
	mj, _ := json.Marshal(m)

	err := deps.KafkaCli.ProduceMessage(kafka.Message{
		Value: mj,
	})

	return err
}
