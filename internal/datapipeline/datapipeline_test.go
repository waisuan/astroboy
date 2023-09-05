package datapipeline

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDataPipeline(t *testing.T) {
	a := assert.New(t)

	go Run()

	time.Sleep(15 * time.Second)

	v, ok := Memory.Load("LatestMessage")
	a.True(ok)
	a.Equal("Hello, Universe!", v)
}
