package main

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/workers"
	"os"
	"os/signal"
)

func main() {
	deps := dependencies.Init()

	w := workers.NewJobEnqueuer(deps.CacheCli)

	w.Pool.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	w.Pool.Stop()
}
