package cmd

import (
	"os"
	"os/signal"
)

func mustGetwd() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return dir
}

func trapSignal(cleanupFunc func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		// block until we've received a signal
		<-c

		// execute any cleanup logic
		if cleanupFunc != nil {
			cleanupFunc()
		}
	}()
}
