package service

import (
	"bufio"
	"fmt"
	"github.com/DevtronLabs/GoToProject/common"
	"github.com/DevtronLabs/GoToProject/internal/constants"
	"github.com/DevtronLabs/GoToProject/pkg/utils"
	"os"
	"strings"
	"sync"
	"time"
)

type ProducerServiceObj struct{}

// Starts the producer
func (service ProducerServiceObj) RunProducer(producer chan<- common.InputMessage) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Producer CLI - Enter messages in the format '<topic>:<message>:<processing_time>:<count>' (Ctrl+C to quit):")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		//Validation
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		msg, err := utils.ParseMessage(input)
		if err != nil {
			fmt.Println("Invalid input format. Please use '<topic>:<message>:<processing_time>:<count>' format.")
			continue
		}

		go func() {
			producer <- msg
		}()
	}
}

// Write the message to the queue file
func (service ProducerServiceObj) WriteToQueueFile(msg common.InputMessage) error {
	queueDir := fmt.Sprintf("%s/%s", constants.QueueDir, msg.Topic)
	if err := os.MkdirAll(queueDir, os.ModePerm); err != nil {
		return err
	}

	filename := fmt.Sprintf("%s/%d.txt", queueDir, time.Now().UnixNano())
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	mutex := sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()

	for i := 0; i < msg.Count; i++ {
		line := fmt.Sprintf("%s:%s:%s\n", msg.Topic, msg.Message, msg.ProcessingTime.String())
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}

	return nil
}
