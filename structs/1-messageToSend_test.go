package structs

import (
	"fmt"
	"testing"
)

func getMessageText(m MessageToSend) string {
	return fmt.Sprintf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
}

func TestMessages(t *testing.T) {
	type testCase struct {
		message     string
		expected    string
		phoneNumber int
	}
	tests := []testCase{
		{
			"Thanks for signing up",
			"Sending message: 'Thanks for signing up' to: 148255510981\n",
			148255510981,
		},
		{
			"Love to have you aboard!",
			"Sending message: 'Love to have you aboard!' to: 148255510982\n",
			148255510982,
		},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{
				"We're so excited to have you",
				"Sending message: 'We're so excited to have you' to: 148255510983\n",
				148255510983,
			},
			{
				"",
				"Sending message: '' to: 148255510984\n",
				148255510984,
			},
			{
				"Hello, World!",
				"Sending message: 'Hello, World!' to: 148255510985\n",
				148255510985,
			},
			{
				"This is a test message",
				"Sending message: 'This is a test message' to: 148255510986\n",
				148255510986,
			},
		}...)
	}

	for _, test := range tests {
		if output := getMessageText(MessageToSend{
			phoneNumber: test.phoneNumber,
			message:     test.message,
		}); output != test.expected {
			t.Errorf(
				"Test Failed: (%v, %v) -> expected: %v actual: %v",
				test.phoneNumber,
				test.message,
				test.expected,
				output,
			)
		} else {
			fmt.Printf("Test Passed: (%v, %v) -> expected: %v actual: %v\n",
				test.phoneNumber,
				test.message,
				test.expected,
				output,
			)
		}
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
