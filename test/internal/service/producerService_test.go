package service

import (
	"fmt"
	"github.com/DevtronLabs/GoToProject/common"
	"github.com/DevtronLabs/GoToProject/internal/service"
	"os"
	"sync"
	"testing"
	"time"
)

type MockMutex struct {
	sync.Mutex
	locked bool
}

func (m *MockMutex) Lock() {
	m.locked = true
}

func (m *MockMutex) Unlock() {
	m.locked = false
}

func TestRunProducer(t *testing.T) {
	mockStdin := "topic1:message1:10s:5\ntopic2:message2:5s:3\n"
	mockProducer := make(chan common.InputMessage, 2)

	// Mock os.Stdin with mockStdin
	r, w, _ := os.Pipe()
	fmt.Fprint(w, mockStdin)
	w.Close()
	oldStdin := os.Stdin
	os.Stdin = r

	// Restore os.Stdin when the test is done
	defer func() {
		os.Stdin = oldStdin
	}()

	// Start the producer
	go service.ProducerServiceObj{}.RunProducer(mockProducer)

	// Wait for some time to allow the producer to process the input
	time.Sleep(100 * time.Millisecond)

	// Check the produced messages
	expectedMessages := []common.InputMessage{
		{
			Topic:          "topic1",
			Message:        "message1",
			ProcessingTime: 10 * time.Second,
			Count:          5,
		},
		{
			Topic:          "topic2",
			Message:        "message2",
			ProcessingTime: 5 * time.Second,
			Count:          3,
		},
	}

	for _, expected := range expectedMessages {
		select {
		case msg := <-mockProducer:
			if msg != expected {
				t.Errorf("unexpected produced message, got: %v, want: %v", msg, expected)
			}
		case <-time.After(100 * time.Millisecond):
			t.Errorf("timeout: no message produced")
		}
	}
}
