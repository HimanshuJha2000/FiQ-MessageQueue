package controllers

import (
	"github.com/DevtronLabs/GoToProject/common"
	"github.com/DevtronLabs/GoToProject/internal/service"
)

type ConsumerInterface interface {
	ProcessMessage(common.InputMessage)
	Run(<-chan common.InputMessage)
}

type ConsumerStruct struct {
	ConsumerService service.ConsumerServiceObj
}

func (consumerObj *ConsumerStruct) ProcessMessage(inputMessage common.InputMessage) {
	consumerObj.ConsumerService.ProcessMessage(inputMessage)
}

func (consumerObj *ConsumerStruct) Run(consumer <-chan common.InputMessage) {
	consumerObj.ConsumerService.RunConsumer(consumer)
}
