package controllers

import (
	"fmt"
	"github.com/MyOrg/FiQ-MessageQueue/common"
	"github.com/MyOrg/FiQ-MessageQueue/internal/service"
)

type ProducerInterface interface {
	ProcessMessage(common.InputMessage)
	Run(chan<- common.InputMessage)
}

type ProducerStruct struct {
	producerService service.ProducerServiceObj
}

func (producerObj *ProducerStruct) ProcessMessage(inputMessage common.InputMessage) {
	err := producerObj.producerService.WriteToQueueFile(inputMessage)
	if err != nil {
		fmt.Print("Error occurred while writing this message to the file ", inputMessage)
		return
	}
}

func (producerObj *ProducerStruct) Run(producer chan<- common.InputMessage) {
	producerObj.producerService.RunProducer(producer)
}
