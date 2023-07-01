package bootstrap

import (
	"github.com/DevtronLabs/GoToProject/common"
	"github.com/DevtronLabs/GoToProject/internal/controllers"
)

func InitConsumer(consumerObj controllers.ConsumerInterface, consumer <-chan common.InputMessage) {
	go consumerObj.Run(consumer)
}

func StartConsuming(consumerObj controllers.ConsumerInterface, consumer chan common.InputMessage, done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
		case inputMsg := <-consumer:
			consumerObj.ProcessMessage(inputMsg)
		}
	}
}

func InitProducer(producerObj controllers.ProducerInterface, producer chan<- common.InputMessage) {
	go producerObj.Run(producer)
}

func StartProducing(producerObj controllers.ProducerInterface, producer chan common.InputMessage, done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
		case inputMsg := <-producer:
			producerObj.ProcessMessage(inputMsg)
		}
	}
}
