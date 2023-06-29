package cmd

import (
	"github.com/DevtronLabs/GoToProject/common"
	"github.com/DevtronLabs/GoToProject/internal/bootstrap"
	"github.com/DevtronLabs/GoToProject/internal/controllers"
	"os"
	"os/signal"
	"syscall"
)

func StartingConsumer() {
	done := make(chan struct{})
	consumer := make(chan common.InputMessage)

	consumerObj := &controllers.ConsumerStruct{}
	bootstrap.InitConsumer(consumerObj, consumer)

	// Wait for termination signal (SIGINT)
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		<-signalChannel
		close(done)
	}()

	bootstrap.StartConsuming(consumerObj, consumer, done)

	<-done
}
