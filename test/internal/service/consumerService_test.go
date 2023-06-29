package service

import (
	"fmt"
	"github.com/DevtronLabs/GoToProject/common"
	service2 "github.com/DevtronLabs/GoToProject/internal/service"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"testing"
	"time"
)

type MockConsumer chan common.InputMessage

func (m MockConsumer) Close() error {
	close(m)
	return nil
}

func (m MockConsumer) Enqueue(msg common.InputMessage) {
	m <- msg
}

func TestRunConsumer(t *testing.T) {
	consumer := make(MockConsumer, 10)
	inputFile, err := ioutil.TempFile("", "mock_input")
	if err != nil {
		t.Fatal("Failed to create mock input file:", err)
	}
	defer os.Remove(inputFile.Name())

	mockInput := "2\n"
	if _, err := inputFile.Write([]byte(mockInput)); err != nil {
		t.Fatal("Failed to write mock input:", err)
	}
	if _, err := inputFile.Seek(0, io.SeekStart); err != nil {
		t.Fatal("Failed to seek to the beginning of the mock input file:", err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = os.NewFile(uintptr(syscall.Stdin), inputFile.Name())

	service := service2.ConsumerServiceObj{}

	go service.RunConsumer(consumer)

	msg1 := common.InputMessage{
		Topic:          "test-topic-1",
		Message:        "test-message-1",
		ProcessingTime: time.Millisecond * 100,
	}
	msg2 := common.InputMessage{
		Topic:          "test-topic-2",
		Message:        "test-message-2",
		ProcessingTime: time.Millisecond * 200,
	}
	consumer.Enqueue(msg1)
	consumer.Enqueue(msg2)

	time.Sleep(time.Millisecond * 500)

	consumer.Close()

	assert.Equal(t, 2, len(consumer), "Unexpected number of processed messages")

}

func TestProcessMessage(t *testing.T) {
	service := service2.ConsumerServiceObj{}

	msg := common.InputMessage{
		Topic:          "test-topic",
		Message:        "test-message",
		ProcessingTime: time.Millisecond * 100,
	}

	startTime := time.Now()
	service.ProcessMessage(msg)
	elapsedTime := time.Since(startTime)

	// Check that the processing time is within an acceptable range
	assert.InDelta(t, float64(elapsedTime), float64(msg.ProcessingTime), float64(10*time.Millisecond), "Unexpected processing time")

}

func TestProcessTopicDir(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := ioutil.TempDir("", "test_topic_dir")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer cleanupTempDir(tempDir)

	// Create some test files inside the temporary directory
	testFiles := []string{"file1.txt", "file2.txt", "file3.txt"}
	for _, filename := range testFiles {
		filePath := filepath.Join(tempDir, filename)
		if err := ioutil.WriteFile(filePath, []byte("test"), 0644); err != nil {
			t.Fatalf("Failed to create test file %s: %v", filename, err)
		}
	}

	// Set up the message channel
	messageCh := make(chan string, len(testFiles))

	// Call the function under test
	service2.ProcessTopicDir(tempDir, messageCh)

	// Verify the messages received
	expectedMessages := []string{"test", "test", "test"}
	for i, expected := range expectedMessages {
		received := <-messageCh
		if received != expected {
			t.Errorf("Mismatched message at index %d. Expected: %s, Received: %s", i, expected, received)
		}
	}

	// Ensure that no extra messages were sent
	select {
	case received := <-messageCh:
		t.Errorf("Received unexpected extra message: %s", received)
	default:
		// No extra messages, which is expected
	}
}

func cleanupTempDir(tempDir string) {
	if err := os.RemoveAll(tempDir); err != nil {
		fmt.Printf("Failed to clean up temporary directory %s: %v\n", tempDir, err)
	}
}
