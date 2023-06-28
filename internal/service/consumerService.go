package service

import (
	"bufio"
	"fmt"
	"github.com/DevtronLabs/GoToProject/common"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ConsumerServiceObj struct{}

func (service ConsumerServiceObj) RunConsumer(consumer <-chan common.InputMessage) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Consumer CLI - Enter the number of concurrent workers:")
	fmt.Print("> ")

	//Validation
	concurrency, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	concurrency = strings.TrimSpace(concurrency)
	if concurrency == "" {
		return
	}

	workers, err := strconv.Atoi(concurrency)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return

	}

	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(id int) {
			defer wg.Done()
			for msg := range consumer {
				service.ProcessMessage(msg)
			}
		}(i + 1)
	}

	wg.Wait()
}

func (service ConsumerServiceObj) ProcessMessage(msg common.InputMessage) {
	fmt.Printf("Processing message (topic: %s): %s\n", msg.Topic, msg.Message)
	time.Sleep(msg.ProcessingTime)
	fmt.Printf("Finished processing message (topic: %s): %s\n", msg.Topic, msg.Message)
}
