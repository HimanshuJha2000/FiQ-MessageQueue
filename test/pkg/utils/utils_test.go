package utils

import (
	"fmt"
	"github.com/DevtronLabs/GoToProject/common"
	"github.com/DevtronLabs/GoToProject/pkg/utils"
	"testing"
	"time"
)

func TestParseMessage(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      common.InputMessage
		expectedError error
	}{
		{
			name:  "ValidInput",
			input: "topic:message:10s:5",
			expected: common.InputMessage{
				Topic:          "topic",
				Message:        "message",
				ProcessingTime: 10 * time.Second,
				Count:          5,
			},
			expectedError: nil,
		},
		{
			name:          "InvalidInputFormat",
			input:         "topic:message:10s",
			expected:      common.InputMessage{},
			expectedError: fmt.Errorf("invalid input format"),
		},
		{
			name:          "InvalidProcessingTime",
			input:         "topic:message:invalid:5",
			expected:      common.InputMessage{},
			expectedError: fmt.Errorf("invalid processing time: time: invalid duration \"invalid\""),
		},
		{
			name:          "InvalidCount",
			input:         "topic:message:10s:invalid",
			expected:      common.InputMessage{},
			expectedError: fmt.Errorf("invalid count: strconv.Atoi: parsing \"invalid\": invalid syntax"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := utils.ParseMessage(test.input)

			// Check the error
			if err != nil {
				if test.expectedError == nil {
					t.Errorf("unexpected error: %v", err)
				} else if err.Error() != test.expectedError.Error() {
					t.Errorf("unexpected error, got: %v, want: %v", err, test.expectedError)
				}
			} else if test.expectedError != nil {
				t.Errorf("expected error: %v, but got nil", test.expectedError)
			}

			// Check the result
			if result != test.expected {
				t.Errorf("unexpected result, got: %v, want: %v", result, test.expected)
			}
		})
	}
}
