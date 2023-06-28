package common

import "time"

type InputMessage struct {
	Topic          string        `json:"topic"`
	Message        string        `json:"message"`
	ProcessingTime time.Duration `json:"processing_time"`
	Count          int           `json:"count"`
}
