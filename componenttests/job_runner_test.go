package componenttests

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/workers"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestJobRunner(t *testing.T) {
	a := assert.New(t)
	os.Setenv("APP_ENV", "test")

	deps := dependencies.Init()

	kafkaMailbox := make(chan string)
	defer close(kafkaMailbox)

	go deps.KafkaCli.ConsumeMessage(kafkaMailbox)

	w := workers.NewJobEnqueuer(deps)
	w.Pool.Start()

	select {
	case kafkaMessage := <-kafkaMailbox:
		a.NotEmpty(kafkaMessage)
	case <-time.After(10 * time.Second):
		a.FailNow("unable to determine outcome of background job")
	}

	w.Pool.Stop()
}
