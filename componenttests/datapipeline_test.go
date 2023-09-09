package componenttests

import (
	"astroboy/internal/cache"
	k "astroboy/internal/kafka"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDataPipeline(t *testing.T) {
	a := assert.New(t)

	kafkaCli := k.NewKafkaCli()

	err := kafkaCli.ProduceMessage(kafka.Message{
		Key:   []byte("Some-Key"),
		Value: []byte("Hello, Universe!"),
	})
	a.Nil(err)

	cacheCli := cache.NewCache()

	found := false
	for i := 0; i < 5; i++ {
		v, err := cacheCli.Get("latest_message")
		require.Nil(t, err)

		if v != "" {
			a.Equal("Hello, Universe!", v)
			found = true
			break
		}

		time.Sleep(time.Duration(i) * time.Second * 2)
	}

	if !found {
		a.Fail("expected to receive a message but got nothing")
	}
}
