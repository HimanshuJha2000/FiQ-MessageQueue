package main

import (
	"github.com/DevtronLabs/GoToProject/common"
	"github.com/DevtronLabs/GoToProject/internal/bootstrap"
	"github.com/DevtronLabs/GoToProject/internal/controllers"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	done := make(chan struct{})
	producer := make(chan common.InputMessage)

	producerObj := &controllers.ProducerStruct{} // Initialize the producerObj
	bootstrap.InitProducer(producerObj, producer)

	// Wait for termination signal (SIGINT)
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		<-signalChannel
		close(done)
	}()

	bootstrap.StartProducing(producerObj, producer, done)

	<-done
}
