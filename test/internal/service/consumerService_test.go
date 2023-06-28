package service

import (
	"github.com/DevtronLabs/GoToProject/common"
	service2 "github.com/DevtronLabs/GoToProject/internal/service"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"os"
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
