package main

import (
	"bufio"
	"fmt"
	cmd2 "github.com/DevtronLabs/GoToProject/cmd/consumer"
	cmd "github.com/DevtronLabs/GoToProject/cmd/producer"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the CLI tools!")
	fmt.Println("Please enter 0 to run as a consumer or 1 to run as a producer: ")

	// Read the user's input
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Convert the input to an integer
	option, err := strconv.Atoi(input[:len(input)-1]) // Removing the newline character
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	// Check the user's choice
	switch option {
	case 0:
		fmt.Println("Running as a consumer...")
		cmd2.StartingConsumer()
	case 1:
		fmt.Println("Running as a producer...")
		cmd.StartingProducer()
	default:
		fmt.Println("Invalid option. Please choose either 0 or 1.")
	}
}
