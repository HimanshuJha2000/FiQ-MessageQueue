package service

import (
	"bufio"
	"fmt"
	"github.com/DevtronLabs/GoToProject/common"
	"github.com/DevtronLabs/GoToProject/internal/constants"
	"io/ioutil"
	"os"
	"path/filepath"
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

	topics, err := ioutil.ReadDir(constants.QueueDir)
	if err != nil {
		fmt.Println("Failed to read the queue directory:", err)
		os.Exit(1)
	}

	// Create a channel to receive messages
	messageCh := make(chan string)

	// Create a wait group for the workers
	var wg sync.WaitGroup

	// Start the workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, messageCh)
		}(i + 1)
	}

	for _, topic := range topics {
		if topic.IsDir() {
			topicDir := filepath.Join(constants.QueueDir, topic.Name())
			ProcessTopicDir(topicDir, messageCh)
		}
	}

	// Close the channel to signal the workers to exit
	close(messageCh)

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("fiq_queue consumer consumption completed")
}

func ProcessTopicDir(topicDir string, messageCh chan<- string) {
	// Read the messages from the topic directory
	files, err := ioutil.ReadDir(topicDir)
	if err != nil {
		fmt.Printf("Failed to read topic directory %s: %v\n", topicDir, err)
		return
	}

	fileWg := sync.WaitGroup{}
	// Process each file and send messages to the channel
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(topicDir, file.Name())
			fileWg.Add(1)
			go processFile(filePath, messageCh, &fileWg)
		}
	}
	// Wait for all files to be processed
	fileWg.Wait()
}

func processFile(filePath string, messageCh chan<- string, wg *sync.WaitGroup) {
	// Read the contents of the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read file %s: %v\n", filePath, err)
		wg.Done()
		return
	}

	// Split the content into individual messages
	messages := strings.Split(string(content), "\n")

	for _, message := range messages {
		if message == "" {
			continue
		}

		// Send the message to the channel
		messageCh <- message
	}

	wg.Wait()

	// Remove the processed file
	err = os.Remove(filePath)
	if err != nil {
		fmt.Printf("Failed to remove file %s: %v\n", filePath, err)
	}

	wg.Done()
}

func worker(id int, messageCh <-chan string) {
	fmt.Printf("Worker %d started\n", id)

	for message := range messageCh {
		parts := strings.Split(message, ":")
		if len(parts) != 3 {
			fmt.Println("Invalid message format:", message)
			continue
		}

		//topic := parts[0]
		messageText := parts[1]
		processingTime, err := time.ParseDuration(parts[2])
		if err != nil {
			fmt.Println("Invalid processing time format:", parts[2])
			continue
		}

		fmt.Printf("Worker %d processing message: %s\n", id, messageText)
		time.Sleep(processingTime)
		fmt.Printf("Worker %d completed message: %s\n", id, messageText)
	}

	fmt.Printf("Worker %d stopped\n", id)
}

func (service ConsumerServiceObj) ProcessMessage(msg common.InputMessage) {
	fmt.Printf("Processing message (topic: %s): %s\n", msg.Topic, msg.Message)
	time.Sleep(msg.ProcessingTime)
	fmt.Printf("Finished processing message (topic: %s): %s\n", msg.Topic, msg.Message)
}
