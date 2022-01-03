package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	_, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGTERM)

	done := make(chan bool, 1)
	go func() {
		<-signalChannel
		done <- true
	}()

	select {
	case <-done:
		cancelFunc()
	}
}
