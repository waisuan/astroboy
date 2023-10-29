package componenttests

import (
	"testing"
)

func TestJobRunner(t *testing.T) {
	//a := assert.New(t)
	//os.Setenv("APP_ENV", "test")
	//
	//deps := dependencies.Init()
	//
	//kafkaMailbox := make(chan string)
	//defer close(kafkaMailbox)
	//
	//go func() {
	//	err := deps.KafkaCli.ConsumeMessage(kafkaMailbox)
	//	if err != nil {
	//		a.FailNow("failed to consume Kafka message: %e", err)
	//	}
	//}()
	//
	//w := workers.NewJobEnqueuer(deps)
	//w.Pool.Start()
	//
	//select {
	//case kafkaMessage := <-kafkaMailbox:
	//	a.NotEmpty(kafkaMessage)
	//case <-time.After(300 * time.Second):
	//	a.FailNow("unable to determine outcome of background job")
	//}
	//
	//w.Pool.Stop()
}
