package cmd

import (
	"github.com/MyOrg/FiQ-MessageQueue/common"
	"github.com/MyOrg/FiQ-MessageQueue/internal/bootstrap"
	"github.com/MyOrg/FiQ-MessageQueue/internal/controllers"
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
