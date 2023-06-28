package utils

import (
	"fmt"
	"github.com/DevtronLabs/GoToProject/common"
	"strconv"
	"strings"
	"time"
)

func ParseMessage(input string) (common.InputMessage, error) {
	parts := strings.Split(input, ":")
	if len(parts) != 4 {
		return common.InputMessage{}, fmt.Errorf("invalid input format")
	}

	topic := strings.TrimSpace(parts[0])
	message := strings.TrimSpace(parts[1])

	processingTime, err := time.ParseDuration(strings.TrimSpace(parts[2]))
	if err != nil {
		return common.InputMessage{}, fmt.Errorf("invalid processing time: %v", err)
	}

	count, err := strconv.Atoi(strings.TrimSpace(parts[3]))
	if err != nil {
		return common.InputMessage{}, fmt.Errorf("invalid count: %v", err)
	}

	return common.InputMessage{
		Topic:          topic,
		Message:        message,
		ProcessingTime: processingTime,
		Count:          count,
	}, nil
}
